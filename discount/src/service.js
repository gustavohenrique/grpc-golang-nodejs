
const DISCOUNTS = {
  birthDate: process.env.BIRTHDATE_DISCOUNT_VALUE || 0.05,
  blackFriday: process.env.BLACKFRIDAY_DISCOUNT_VALUE || 0.1
};

class Service {
  constructor(repository) {
    this.repository = repository;
  }

  applyDiscount(params) {
    const { customer, product } = params;

    console.log('I will try to apply discount for', JSON.stringify(customer));

    let discountInDecimal;
    if (this._isBirthDate(customer)) {
      discountInDecimal = DISCOUNTS.birthDate;
      console.log('Happy birthday to you =). You win discount:', discountInDecimal);
    }

    if (this._isBlackFriday()) {
      discountInDecimal = DISCOUNTS.blackFriday;
      console.log('Enjoy our insane blackfriday. You win discount:', discountInDecimal);
    }

    try {
      const discount = this.repository.findDiscountByProductID(product.id)
      console.log('discount discountInDecimal', discount)
      const maxDiscountValue = discount.max_discount;
      if (discountInDecimal > maxDiscountValue) {
        discountInDecimal = maxDiscountValue;
        console.log('Opps. Looks like you have a big discount. Sorry, but I will decrease your discount to:', discountInDecimal);
      }
      return this._calculateDiscount(product.priceInCents, discountInDecimal);
    } catch (err) {
      console.log('Sorry, I cannot give you a discount for this product.', err.name);
      return {
        pct: 0,
        valueInCents: 0
      };
    }
  }

  _calculateDiscount(price, discountInDecimal) {
    price = parseInt(price);
    const newPrice = Math.round(price * discountInDecimal);
    return {
      pct: discountInDecimal,
      valueInCents: price - newPrice
    };
  }

  _isBirthDate(customer) {
    const dt = new Date().toISOString().split('T')[0];
    const [_, month, day] = dt.split('-');
    const birthDate = customer.birthDate.split('T')[0].split('-');
    return `${month}-${day}` === `${birthDate[1]}-${birthDate[2]}`;
  }

  _isBlackFriday() {
    const dt = new Date().toISOString().split('T')[0];
    const [_, month, day] = dt.split('-');
    const blackFridayDate = process.env.BLACKFRIDAY_DATE || '11-25';
    return `${month}-${day}` === blackFridayDate;
  }
}

module.exports = Service;