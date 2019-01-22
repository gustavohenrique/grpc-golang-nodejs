const messages = require('../ecommerce_pb');

const Repository = require('./repository');
const Service = require('./service');


const databaseURL = process.env.DATABASE_URL;
const repository = new Repository(databaseURL);
repository.connect();
const discountService = new Service(repository);

module.exports = {
  applyDiscount (call, callback) {
    const customer = call.request.getCustomer().toObject();
    const product = call.request.getProduct().toObject();

    console.log('Asking our robots to calculate your discount...', product.title);

    const discount = discountService.applyDiscount({ customer, product });

    console.log('Congrats! Our robots give to you the discount:', JSON.stringify(discount));
    
    const discountValue = new messages.DiscountValue();
    discountValue.setPct(discount.pct);
    discountValue.setValueInCents(discount.valueInCents);

    const productWithDiscount = new messages.Product();
    productWithDiscount.setId(product.id);
    productWithDiscount.setTitle(product.title)
    productWithDiscount.setDescription(product.description)
    productWithDiscount.setPriceInCents(product.priceInCents)
    productWithDiscount.setDiscountValue(discountValue);

    const response = new messages.DiscountResponse();
    response.setProduct(productWithDiscount);

    console.log('Enjoy your discount.');

    callback(null, response);
  }
}