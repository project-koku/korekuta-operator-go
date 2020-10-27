/*
Copyright 2020 Red Hat, Inc.
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.
You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package packaging

import (
	"archive/tar"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	uuidv4 "github.com/delaemon/go-uuidv4"
	"github.com/go-logr/logr"
	costmgmtv1alpha1 "github.com/project-koku/korekuta-operator-go/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Define the global variables
var defaultMaxSize int64 = 100
var megaByte int64 = 1024 * 1024
var maxBytes int64 = defaultMaxSize * megaByte

// the csv module doesn't expose the bytes-offset of the
// underlying file object.
// instead, the script estimates the size of the data as VARIANCE percent larger than a
// naïve string concatenation of the CSV fields to cover the overhead of quoting
// and delimiters. This gets close enough for now.
// VARIANCE := 0.03
var variance float64 = 0.03

// if we're creating more than 1k files, something is probably wrong.
var maxSplits int64 = 1000

// define the manifest template
type Manifest struct {
	UUID      string   `json:"uuid"`
	ClusterID string   `json:"cluster_id"`
	Version   string   `json:"version"`
	Date      string   `json:"date"`
	Files     []string `json:"files"`
}

func BuildLocalCSVFileList(stagingDirectory string) []string {
	var csvList []string
	fileList, err := ioutil.ReadDir(stagingDirectory)
	if err != nil {
		fmt.Println("could not read dir")
		// log.Error(err, "Could not read the directory")
	}
	for _, file := range fileList {
		if strings.Contains(file.Name(), ".csv") {
			csvList = append(csvList, stagingDirectory+"/"+file.Name())
		}
	}
	return csvList
}

func NeedSplit(filepath string) bool {
	var totalSize int64 = 0
	maxBytes := defaultMaxSize * megaByte
	fileList, err := ioutil.ReadDir(filepath)
	if err != nil {
		fmt.Println("could not read dir")
	}
	for _, file := range fileList {
		info, err := os.Stat(filepath + "/" + file.Name())
		if err != nil {
			return false
		}
		fileSize := info.Size()
		totalSize += fileSize
		if fileSize >= maxBytes || totalSize >= maxBytes {
			return true
		}
	}
	return false
}

func RenderManifest(logger logr.Logger, archiveFiles []string, cost *costmgmtv1alpha1.CostManagement, filepath string) (string, string) {
	log := logger.WithValues("costmanagement", "RenderManifest")
	// setup the manifest
	manifestUUID, _ := uuidv4.Generate()
	manifestDate := metav1.Now()
	var manifestFiles []string
	for idx := range archiveFiles {
		uploadName := manifestUUID + "_openshift_usage_report." + strconv.Itoa(idx) + ".csv"
		manifestFiles = append(manifestFiles, uploadName)
	}
	fileManifest := Manifest{
		UUID:      manifestUUID,
		ClusterID: cost.Status.ClusterID,
		Version:   cost.Status.OperatorCommit,
		Date:      manifestDate.UTC().Format("2006-01-02 15:04:05"),
		Files:     manifestFiles,
	}
	manifestFileName := filepath + "/manifest.json"
	// write the manifest file
	file, _ := json.MarshalIndent(fileManifest, "", " ")
	_ = ioutil.WriteFile(manifestFileName, file, 0644)
	// return the manifest file/uuid
	log.Info("Generated manifest file", "manifest", manifestFileName)
	return manifestFileName, manifestUUID
}

func addFileToTarWriter(logger logr.Logger, uploadName, filePath string, tarWriter *tar.Writer) error {
	log := logger.WithValues("costmanagement", "addFileToTarWriter")
	log.Info("Adding file to tar.gz", "file", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    uploadName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}

func WriteTarball(logger logr.Logger, tarFileName, manifestFileName, manifestUUID string, archiveFiles []string, fileNum ...int) string {
	index := 0
	if len(fileNum) > 0 {
		index = fileNum[0]
	}
	if len(archiveFiles) <= 0 {
		return ""
	}
	// create the tarfile
	tarFile, err := os.Create(tarFileName)
	if err != nil {
		fmt.Println("Error!")
	}
	defer tarFile.Close()

	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tw := tar.NewWriter(gzipWriter)
	defer tw.Close()

	// add the files to the tarFile
	for idx, fileName := range archiveFiles {
		if index != 0 {
			idx = index
		}
		fmt.Println(fileName)
		if strings.Contains(fileName, ".csv") {
			uploadName := manifestUUID + "_openshift_usage_report." + strconv.Itoa(idx) + ".csv"
			fmt.Println(uploadName)
			err := addFileToTarWriter(logger, uploadName, fileName, tw)
			if err != nil {
				fmt.Println(err)
				return ""
			}
		}
	}
	addFileToTarWriter(logger, "manifest.json", manifestFileName, tw)

	return tarFileName

}

func WritePart(logger logr.Logger, fileName string, csvReader *csv.Reader, csvHeader []string, num int64) (string, bool, error) {
	log := logger.WithValues("costmanagement", "WritePart")
	fileNamePart := strings.TrimSuffix(fileName, ".csv")
	sizeEstimate := 0
	splitFileName := fileNamePart + strconv.FormatInt(num, 10) + ".csv"
	log.Info("Creating file ", "file", splitFileName)
	splitFile, err := os.Create(splitFileName)
	if err != nil {
		return "", false, fmt.Errorf("WritePart: error creating file: %v", err)
	}
	// Create the csv writer
	writer := csv.NewWriter(splitFile)
	// Preserve the header
	writer.Write(csvHeader)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			writer.Flush()
			return splitFileName, true, nil
		}
		writer.Write(row)
		rowLen := len(strings.Join(row, ","))
		rowSize := rowLen + int(float64(rowLen)*variance)
		sizeEstimate += rowSize
		if sizeEstimate >= int(maxBytes) {
			writer.Flush()
			return splitFileName, false, nil
		}
	}
}

func SplitFiles(logger logr.Logger, filePath string) error {
	fileList, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println("could not read dir")
	}
	for _, file := range fileList {
		absPath := filePath + "/" + file.Name()
		info, err := os.Stat(absPath)
		if err != nil {
			return fmt.Errorf("SplitFiles: error getting fileInfo: %v", err)
		}
		fileSize := info.Size()
		if fileSize >= maxBytes {
			var splitFiles []string
			// var csvHeader string
			// open the file
			csvFile, err := os.Open(absPath)
			if err != nil {
				return fmt.Errorf("SplitFiles: error reading file: %v", err)
			}
			csvReader := csv.NewReader(csvFile)
			csvHeader, err := csvReader.Read()
			var part int64 = 1
			for {
				newFile, eof, err := WritePart(logger, absPath, csvReader, csvHeader, part)
				if err != nil {
					return fmt.Errorf("SplitFiles: %v", err)
				}
				splitFiles = append(splitFiles, newFile)
				part++
				if eof || part >= maxSplits {
					break
				}
			}
			os.Remove(absPath)
			fmt.Println(splitFiles)
		}
	}
	return nil
}

func Split(logger logr.Logger, filePath string, cost *costmgmtv1alpha1.CostManagement) error {
	log := logger.WithValues("costmanagement", "Split")
	var outFiles []string
	log.Info("Checking to see if the report files need to be split")
	needSplit := NeedSplit(filePath)
	if needSplit {
		log.Info("Report files exceed the max size. Splitting files")
		if err := SplitFiles(logger, filePath); err != nil {
			return fmt.Errorf("Split: %v", err)
		}
		tarPath := filePath + "/../"
		tarFileTmpl := "cost-mgmt"
		fileList := BuildLocalCSVFileList(filePath)
		manifestFileName, manifestUUID := RenderManifest(logger, fileList, cost, filePath)
		for idx, fileName := range fileList {
			if strings.Contains(fileName, ".csv") {
				fileList = []string{fileName}
				tarFileName := tarPath + tarFileTmpl + strconv.Itoa(idx) + ".tar.gz"
				log.Info("Generating tar.gz", "tarFile", tarFileName)
				outputTar := WriteTarball(logger, tarFileName, manifestFileName, manifestUUID, fileList, idx)
				if outputTar != "" {
					outFiles = append(outFiles, outputTar)
				}
			}
		}
	} else {
		tarFileName := filePath + "/../cost-mgmt.tar.gz"
		log.Info("Report files do not require split, generating tar.gz", "tarFile", tarFileName)
		fileList := BuildLocalCSVFileList(filePath)
		if len(fileList) > 0 {
			manifestFileName, manifestUUID := RenderManifest(logger, fileList, cost, filePath)
			outputTar := WriteTarball(logger, tarFileName, manifestFileName, manifestUUID, fileList)
			if outputTar != "" {
				outFiles = append(outFiles, outputTar)
			}
		}
	}
	log.Info("Created the following files for upload: ", "files", outFiles)
	return nil
}