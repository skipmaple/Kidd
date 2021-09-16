package main

import "time"

// Sign In and Sign Out
// page 38
type oneDayRecord struct {
	count       int
	signRecords []signRecord
}

type signRecord struct {
	id      string
	signIn  string
	signOut string
}

func signInAndSignOut(someDayRecord oneDayRecord) (firstID, lastID string) {
	temp := signRecord{
		signIn:  "23:59:59",
		signOut: "00:00:00",
	}

	for _, signRecord := range someDayRecord.signRecords {
		tSignIn, _ := time.Parse("15:04:05", signRecord.signIn)
		tSignOut, _ := time.Parse("15:04:05", signRecord.signOut)

		tTempSignIn, _ := time.Parse("15:04:05", temp.signIn)
		tTempSignOut, _ := time.Parse("15:04:05", temp.signOut)

		if tSignIn.Before(tTempSignIn) {
			temp.signIn = signRecord.signIn
			firstID = signRecord.id
		}

		if tSignOut.After(tTempSignOut) {
			temp.signOut = signRecord.signOut
			lastID = signRecord.id
		}
	}

	return firstID, lastID
}

// 时间比较demo

//func main() {
//	signIn := "23:59:59"
//	signOut := "00:00:00"
//	t1, _ := time.Parse("15:04:05", signIn)
//	t2, _ := time.Parse("15:04:05", signOut)
//	fmt.Println(t1.Before(t2))
//}
