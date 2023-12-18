package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMealPlan(t *testing.T) {
	for i, tt := range []struct {
		inPrices         []int
		outMoneyRequired int
		outCouponsLeft   int
		outCouponDays    []int
	}{{
		inPrices:         []int{35, 40, 101, 59, 63},
		outMoneyRequired: 235,
		outCouponsLeft:   0,
		outCouponDays:    []int{4},
	}, {
		inPrices:         []int{101, 102, 103, 104, 105},
		outMoneyRequired: 306,
		outCouponsLeft:   1,
		outCouponDays:    []int{3, 4},
	}, {
		inPrices:         []int{101},
		outMoneyRequired: 101,
		outCouponsLeft:   1,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{100},
		outMoneyRequired: 100,
		outCouponsLeft:   0,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{101, 102, 103, 104, 1, 2, 3, 4, 99, 101, 5, 6, 7, 8, 9},
		outMoneyRequired: 351,
		outCouponsLeft:   0,
		outCouponDays:    []int{3, 8, 9},
	}, {
		inPrices:         []int{},
		outMoneyRequired: 0,
		outCouponsLeft:   0,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{0, 0, 0, 101, 0, 0, 101},
		outMoneyRequired: 101,
		outCouponsLeft:   0,
		outCouponDays:    []int{6},
	}, {
		inPrices:         []int{0, 0, 100, 101, 0, 0, 0},
		outMoneyRequired: 201,
		outCouponsLeft:   1,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{101, 101, 101, 101, 101, 101, 101, 101},
		outMoneyRequired: 404,
		outCouponsLeft:   0,
		outCouponDays:    []int{4, 5, 6, 7},
	}, {
		inPrices:         []int{0, 0, 0, 0},
		outMoneyRequired: 0,
		outCouponsLeft:   0,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{0, 0, 0, 0},
		outMoneyRequired: 0,
		outCouponsLeft:   0,
		outCouponDays:    nil,
	}, {
		inPrices:         []int{101, 100, 101, 101, 99, 101, 101, 101, 1, 101, 101, 101},
		outMoneyRequired: 605,
		outCouponsLeft:   0,
		outCouponDays:    []int{1, 7, 9, 10, 11},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mp := createMealPlan(tt.inPrices)
			assert.Equal(t, tt.outMoneyRequired, mp.moneyRequired)
			assert.Equal(t, tt.outCouponsLeft, mp.couponsLeft)
			assert.Equal(t, tt.outCouponDays, mp.couponDays)
		})
	}
}

func FuzzCreateMealPlan(f *testing.F) {
	for _, seedDays := range [][]byte{
		{35, 40, 101, 59, 63},
		{101, 102, 103, 104, 1, 2, 3, 4, 99, 101, 5, 6, 7, 8, 9},
		{},
		{0},
		{0, 0, 0},
		{101, 0, 0},
		{101, 101, 101, 101, 101, 101, 101, 101},
	} {
		f.Add(seedDays)
	}
	f.Fuzz(func(t *testing.T, pricesAsBytes []byte) {
		prices := make([]int, len(pricesAsBytes))
		for i, priceAsByte := range pricesAsBytes {
			prices[i] = int(priceAsByte)
		}
		mp := createMealPlan(prices)
		moneyRequired := 0
	daysLoop:
		for i, price := range prices {
			for _, couponDay := range mp.couponDays {
				if i == couponDay {
					continue daysLoop
				}
			}
			moneyRequired += price
		}
		assert.Equal(t, moneyRequired, mp.moneyRequired)
	})
}
