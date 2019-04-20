package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// MatchustomerHouses
func MatchustomerHouses() {

	start := time.Now()

	var err error

	RowCustomer, err = DB.Query("select * from " + TableCustomer)
	CheckErr(err)

	_ = DB.QueryRow(`select count(*) from ` + TableCustomer).Scan(&NumOfRowCustomer)

	for RowCustomer.Next() {
		var mrtg Mortgage

		err = RowCustomer.Scan(&mrtg.CID, &mrtg.FirstName, &mrtg.LastName, &mrtg.CreditScore,
			&mrtg.Salary, &mrtg.DownPayment, &mrtg.HouseID)

		CheckErr(err)

		query1 := "select * from " + TableHouse + " where id=" +
			strconv.FormatUint(uint64(mrtg.HouseID), 10)

		rowHouse := DB.QueryRow(query1)

		err = rowHouse.Scan(&mrtg.HID, &mrtg.Price, &mrtg.MinDownPayment,
			&mrtg.PropertyTax, &mrtg.MaintenanceCost)
		CheckErr(err)

		Mrtgs = append(Mrtgs, mrtg)
	}

	WaitG.Add(NumOfRowCustomer)
	for i, mrtg := range Mrtgs {
		go checkMortgageStatus(mrtg, i)
	}
	WaitG.Wait()

	elapsed := time.Since(start)

	fmt.Println("=======================================================")
	fmt.Printf("== MatchCustomerHouse() - time elapsed %s \n", elapsed)
	fmt.Println("=======================================================\n")
}

func checkMortgageStatus(mrtg Mortgage, idx int) {

	Mu.Lock()
	delay := time.Duration(rand.Intn(700) + 300)
	time.Sleep(delay * time.Millisecond)

	approved := false

	if (mrtg.DownPayment < mrtg.MinDownPayment) || (mrtg.CreditScore < 650) {
		approved = false
	}
	yearlyMortgage := (mrtg.Price - mrtg.DownPayment) * 0.0612
	yearlyExpence := yearlyMortgage + mrtg.PropertyTax + mrtg.MaintenanceCost

	if yearlyExpence <= (mrtg.Salary * 0.32) {
		approved = true
	}

	// fmt.Printf("%2d %-10s %-10s %4d %12.2f %12.2f %3d\n",
	// 	mrtg.CID, mrtg.FirstName, mrtg.LastName,
	// 	mrtg.CreditScore, mrtg.Salary, mrtg.DownPayment, mrtg.HouseID)

	// fmt.Printf("%2d %12.2f %12.2f %12.2f %12.2f\n", mrtg.HID, mrtg.Price,
	// 	mrtg.MinDownPayment, mrtg.PropertyTax, mrtg.MaintenanceCost)

	fmt.Println("approved ", approved)
	if approved {
		// fmt.Printf("** Mortgage Application Approved [process time=%s]!\n\n", delay)
		ApproveIdx = append(ApproveIdx, idx)
	} else {
		// fmt.Printf("** Mortgage Application Rejected [process time=%s]!\n\n", delay)
		RejectedIdx = append(RejectedIdx, idx)
	}
	Mu.Unlock()

	WaitG.Done()
}
