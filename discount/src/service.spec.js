const expect = require('chai').expect;
const sinon = require('sinon');
const MockDate = require('mockdate');

const Service = require('./service');


describe('Service', function() {

  let service;

  const Daenerys = {
    id: 2,
    firstName: 'Daenerys',
    lastName: 'Targaryen',
    birthDate: '1986-10-23'
  };

  const Snes = {
    id: 1,
    title: 'Super NES Classic',
    description: 'The Super NES Classic Edition',
    priceInCents: 9149
  };

  const stub = sinon.stub().returns({
    max_discount: '0.1'
  });

  beforeEach(() => {
    const repository = { findDiscountByProductID: stub };
    service = new Service(repository);
  });

  afterEach(() => {
    MockDate.reset();
  })

  describe('#applyDiscount', () => {
    it('should return 5% off when customer birth date is today', async () => {
      MockDate.set('2018-10-23T10:00:00Z');
      const discount = await service.applyDiscount({
          customer: Daenerys,
          product: Snes
      });
      expect(stub.called).to.be.true;
      expect(discount.pct).to.equal(0.05);
      expect(discount.valueInCents).to.equal(8692);
    });

    it('should return 10% off on blackfriday', async () => {
      MockDate.set('2018-11-25T10:00:00Z');
      const discount = await service.applyDiscount({
          customer: Daenerys,
          product: Snes
      });
      expect(stub.called).to.be.true;
      expect(discount.pct).to.equal(0.1);
      expect(discount.valueInCents).to.equal(8234);
    });

    it('should dont apply discount when product is not found', async () => {
      const findDiscountByProductID = sinon.stub().throws('Not found');
      const repository = { findDiscountByProductID };
      const service = new Service(repository);
      const discount = await service.applyDiscount({
          customer: Daenerys,
          product: Snes
      });
      expect(stub.called).to.be.true;
      expect(discount.pct).to.equal(0);
      expect(discount.valueInCents).to.equal(0);
    });

  });

});