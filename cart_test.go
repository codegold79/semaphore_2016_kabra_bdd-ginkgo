package cart_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/codegold79/semaphore_2016_kabra_bdd-ginkgo"
)

var _ = Describe("Shopping cart", func() {
	itemA := Item{ID: "itemA", Name: "Item A", Price: 10.20, Qty: 0}
	itemB := Item{ID: "itemB", Name: "Item B", Price: 7.66, Qty: 0}

	// Contrived example of using BeforeEach as a way to reset the cart variable.
	var cart Cart

	BeforeEach(func() {
		cart = Cart{}
		fmt.Println("---->", cart)
	})

	Context("initially", func() {
		It("has 0 items, units, and total amount", func() {
			Expect(cart.TotalUniqueItems()).Should(BeZero())
			Expect(cart.TotalUnits()).Should(BeZero())
			Expect(cart.TotalAmount()).Should(BeZero())
		})
	})

	Context("when a new item is added", func() {
		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnits()
		originalAmount := cart.TotalAmount()

		Context("the shopping cart", func() {
			It("has 1 more unique item, unit, amount than it had earlier", func() {
				cart.AddItem(itemA)
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount + 1))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount + 1))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("when an existing item is added", func() {
		Context("the shopping cart", func() {			
			It("has the same number of unique items/units/price as earlier", func() {
				cart.AddItem(itemA)
				originalItemCount := cart.TotalUniqueItems()
				originalUnitCount := cart.TotalUnits()
				originalAmount := cart.TotalAmount()
				cart.AddItem(itemA)

				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount + 1))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("that has 0 unit of item A", func() {		
		Context("removing item A", func() {
			It("should not change the number of items/units/price", func() {
				cart.AddItem(itemB) // just to mimic the existence other items
				cart.AddItem(itemB) // just to mimic the existence other items
				originalItemCount := cart.TotalUniqueItems()
				originalUnitCount := cart.TotalUnits()
				originalAmount := cart.TotalAmount()
				cart.RemoveItem(itemA.ID, 1)

				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount))
			})
		})
	})

	Context("that has 1 unit of item A", func() {
		Context("removing 1 unit item A", func() {
			It("should reduce the number of items/units/prices by 1", func() {
				cart.AddItem(itemB) // just to mimic the existence other items
				cart.AddItem(itemB) // just to mimic the existence other items
				cart.AddItem(itemA)
				originalItemCount := cart.TotalUniqueItems()
				originalUnitCount := cart.TotalUnits()
				originalAmount := cart.TotalAmount()
				cart.RemoveItem(itemA.ID, 1)

				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 1))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})
	})

	Context("that has 2 units of item A", func() {
		Context("removing 1 unit of item A", func() {
			It("should not reduce the number of items/units/price", func() {
				cart.AddItem(itemB) // just to mimic the existence other items
				cart.AddItem(itemB) // just to mimic the existence other items
				//Reset the cart with 2 units of item A
				cart.AddItem(itemA)
				cart.AddItem(itemA)
				originalItemCount := cart.TotalUniqueItems()
				originalUnitCount := cart.TotalUnits()
				originalAmount := cart.TotalAmount()
				cart.RemoveItem(itemA.ID, 1)

				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 1))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})

		Context("removing 2 units of item A", func() {		
			It("should reduce the number of items/units/price by 1 or twice", func() {
				cart.AddItem(itemB) // just to mimic the existence other items
				cart.AddItem(itemB) // just to mimic the existence other items
				//Reset the cart with 2 units of item A
				cart.AddItem(itemA)
				cart.AddItem(itemA)
				originalItemCount := cart.TotalUniqueItems()
				originalUnitCount := cart.TotalUnits()
				originalAmount := cart.TotalAmount()
				cart.RemoveItem(itemA.ID, 2)

				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 2))
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - 2*itemA.Price))
			})
		})
	})
})
