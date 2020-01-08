package main

import "github.com/sammarth-kapse/FileDownloadManager/repository"

//getDownloadID returns string
// get the dowmload id of a dowmload request
func getDownloadID(downloadRequest repository.DownloadRequest) string {

	downloadInformation := repository.New(downloadRequest)

	downloadInformation.Download()

	return downloadInformation.ID
}

//getDownloadInfomationID
//get the information of a download from itd dowmload id
func getDownloadInformationByID(id string) (*repository.DownloadInformation, bool) {

	return repository.GetDownloadInformationByID(id)
}
