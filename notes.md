# Notes

When I reset cart variable to an empty struct in BeforeEach(), it looked like it got called before every `It()`. That caused many tests to fail because cart should be reset for every outer-most `Context()`, not for each `It()`.

To fix the tests, I grouped all the `It`s in each `Context()` into one `It()`. I went from 18 failed tests down to 6 failures.

The problem now was,

- `cart.AddItem(itemA)` is happening outside `It()`. 
- `BeforeEach()` is reseting `cart.AddItem()` and setting `cart` back to an empty struct.
- When `It()` is called, it does not have the correct number of items.

So, to fix the remaining tests, I put `cart.AddItem()` inside the `It()`.
