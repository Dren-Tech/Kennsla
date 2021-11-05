import { assert } from 'assertthat';

describe('Logger', (): void => {
  it('is true.', async (): Promise<void> => {
    assert.that(true).is.equalTo(true);
  });
});
