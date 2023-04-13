const add = require("./add")
describe('add', () => {
    it('should add 1 and 1 to be equal 2', () => {
        expect(add(1,1)).toEqual(2)
        //1+1 = 2
    });
    it('should add 1 and 1 to be equal 2', () => {
        expect(add(2,3)).toEqual(5)
        //1+1 = 2
    });
});