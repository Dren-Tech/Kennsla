import { assert } from 'assertthat';

describe('Server', (): void => {
  it('is true.', async (): Promise<void> => {
    assert.that(true).is.equalTo(true);
  });
});
