const { Client } = require('pg');


class Repository {
	constructor(databaseURL) {
    const connectionString = databaseURL;
		this.conn = new Client({ connectionString });
  }

  connect() {
    this.conn.connect();
  }

  disconnect() {
    this.conn.end();
  }

  async runQuery(query, params = null) {
    const res = await this.conn.query(query, params);
    return res;
  }

  async findDiscountByProductID(id) {
    const res = await this.runQuery('SELECT * FROM discounts WHERE product_id = $1', [id]);
    const result = res.rows[0];
    if (! result) {
      throw 'Not found';
    }
    return result;
  }
}

module.exports = Repository;