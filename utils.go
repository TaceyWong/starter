/*
 * @Description: 填写描述
 * @Author: WangXinyong/TaceyWong
 * @Date: 2020-09-25 13:51:46
 * @LastEditors: WangXinyong/TaceyWong
 * @LastEditTime: 2020-09-25 15:03:57
 * @FilePath: /starter/utils.go
 */

package starter

import (
	"fmt"
	"log"
	"os"
)

func RmTree(dirPath string) {
	os.RemoveAll(dirPath)
}

func MakeSurePathExist(path string) bool {
	log.Printf("Making sure path exists:%s", path)
	if os.MkdirAll(path, os.ModePerm) != nil {
		log.Printf("%s has existed", path)
		return false
	}
	log.Printf("Created directory at: %s", path)
	return true
}

func WorkIn(dirname string) {

}

// MakeExecutable Make `scriptPath` executable
//
// param scriptPath: The file to change mode
func MakeExecutable(scriptPath string) {
	// stats, err := os.Stat(scriptPath)
	// if err != nil {
	// log.Fatal(err)
	// }
	// os.Chmod(scriptPath, stats.Mode()|filemode.Executable) //100755
}

// PromptAndDelete Ask user if it's okay to delete the previously-downloaded file/directory.
//
// If yes, delete it. If no, checks to see if the old version should be
// reused. If yes, it's reused; otherwise, Cookiecutter exits.
//
// :param path: Previously downloaded zipfile.
// :param noInput: Suppress prompt to delete repo and just delete it.
// :return: True if the content was deleted
func PromptAndDelete(path string, noInput bool) bool {
	okToDelete := false
	if noInput {
		okToDelete = true
	} else {
		question := fmt.Sprintf("You've downloaded %s before. Is it okay to delete and re-download it?", path)
		ReadUserYesNo(&okToDelete, question, true)
	}
	if okToDelete {
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
		if info.IsDir() {
			RmTree(path)
		} else {
			os.Remove(path)
		}

	} else {
		okToReuse := false
		question := "Do you want to re-use the existing version?"
		ReadUserYesNo(&okToReuse, question, true)
		if okToReuse {
			return false
		}
		os.Exit(0)
	}
	return false
}
