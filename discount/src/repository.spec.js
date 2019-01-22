const expect = require('chai').expect;

const Repository = require('./repository');

describe('Repository', function() {

  let repository;

  beforeEach(async () => {
    const databaseURL = process.env.test_database_url || 'postgres://postgres:root@127.0.0.1:5432/test';
    repository = new Repository(databaseURL);
    await repository.connect();
    await repository.runQuery(`
      DROP TABLE IF EXISTS discounts;
      CREATE TABLE discounts (product_id INT, max_discount DECIMAL);
      INSERT INTO discounts VALUES (1, 0.1);
      INSERT INTO discounts VALUES (2, 0.05);
    `);
  });

  afterEach(() => {
    repository.disconnect();
  })

  describe('#findDiscountByProductID', () => {
    it('should return 0.1 as the max discount for the product with ID 1', async () => {
      const discount = await repository.findDiscountByProductID(1);
      expect(discount.max_discount).to.equal('0.1');
    });

    it('should return an error when no discount was found for the product', async () => {
      try {
        await repository.findDiscountByProductID(999);
      } catch(err) {
        expect(err).exist;
      }
    });
  });

});