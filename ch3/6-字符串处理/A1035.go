package main

import "bytes"

// Password
// page 77

type account struct {
	id  string
	pwd string
}

func replacePassword(n int, accounts []account) (count int, resAccount []account) {
	//replaceMap := map[byte]byte{
	//	'1': '@',
	//	'0': '%',
	//	'O': 'o',
	//	'l': 'L',
	//}

	for i := 0; i < len(accounts); i++ {
		if bytes.ContainsAny([]byte{'1', '0', 'O', 'l'}, accounts[i].pwd) {
			tempPwd := bytes.ReplaceAll([]byte(accounts[i].pwd), []byte{'1'}, []byte{'@'})
			tempPwd = bytes.ReplaceAll(tempPwd, []byte{'0'}, []byte{'%'})
			tempPwd = bytes.ReplaceAll(tempPwd, []byte{'O'}, []byte{'o'})
			tempPwd = bytes.ReplaceAll(tempPwd, []byte{'l'}, []byte{'L'})
			count++
			resAccount = append(resAccount, account{
				id:  accounts[i].id,
				pwd: string(tempPwd),
			})
		}
	}

	return count, resAccount
}
