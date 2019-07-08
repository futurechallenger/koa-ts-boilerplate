import { hello } from '../index';

describe('test index file', () => {
  it('hello', () => {
    expect(hello('kkk')).toEqual('hello kkk');
  });
});
