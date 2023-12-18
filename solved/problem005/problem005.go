// 5. Cafe.
// https://coderun.yandex.ru/problem/cafe?currentPage=1&pageSize=10&rowNumber=5
package main

import "fmt"

const couponThreshold = 100

func main() {
	var n int
	fmt.Scan(&n)
	prices := make([]int, n)
	for i := range prices {
		fmt.Scan(&prices[i])
	}

	mp := createMealPlan(prices)

	fmt.Println(mp.moneyRequired)
	fmt.Println(mp.couponsLeft, len(mp.couponDays))
	for _, cd := range mp.couponDays {
		fmt.Println(cd + 1)
	}
}

func createMealPlan(prices []int) mealPlan {
	mealPlans := make([][]mealPlan, len(prices))
	mp := createMealPlanRec(mealPlans, prices, 0, 0)

	//for dayIdx := range mealPlans {
	//	fmt.Printf("%d: ", dayIdx)
	//	for _, mp := range mealPlans[dayIdx] {
	//		fmt.Printf("{%d %d %v} ", mp.moneyRequired, mp.couponsLeft, mp.couponDays)
	//	}
	//	fmt.Printf("\n")
	//}

	return mp
}

func createMealPlanRec(mealPlans [][]mealPlan, prices []int, dayIdx, couponCount int) mealPlan {
	if dayIdx == len(prices) {
		return mealPlan{couponsLeft: couponCount}
	}
	for _, mp := range mealPlans[dayIdx] {
		if mp.couponsGiven == couponCount {
			return mp
		}
	}

	maybeIncCouponCount := couponCount
	if prices[dayIdx] > couponThreshold {
		maybeIncCouponCount += 1
	}
	mealPlanNoCoupon := createMealPlanRec(mealPlans, prices, dayIdx+1, maybeIncCouponCount)
	mealPlanNoCoupon.moneyRequired += prices[dayIdx]
	mealPlanNoCoupon.couponsGiven = couponCount

	if couponCount > 0 {
		mealPlanWithCoupon := createMealPlanRec(mealPlans, prices, dayIdx+1, couponCount-1)
		if mealPlanWithCoupon.moneyRequired < mealPlanNoCoupon.moneyRequired {
			couponDays := make([]int, len(mealPlanWithCoupon.couponDays)+1)
			copy(couponDays[1:], mealPlanWithCoupon.couponDays)
			couponDays[0] = dayIdx
			mealPlanWithCoupon.couponDays = couponDays
			mealPlanWithCoupon.couponsGiven = couponCount
			mealPlans[dayIdx] = append(mealPlans[dayIdx], mealPlanWithCoupon)
			return mealPlanWithCoupon
		}
	}

	mealPlans[dayIdx] = append(mealPlans[dayIdx], mealPlanNoCoupon)
	return mealPlanNoCoupon
}

type mealPlan struct {
	moneyRequired int
	couponsGiven  int
	couponsLeft   int
	couponDays    []int
}
