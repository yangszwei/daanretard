const ERROR_CODE = {
    SUCCESS: 0,
    PENDING: 1,
    INVALID_INPUT: 2,
    INVALID_TARGET: 3,
    ALREADY_EXIST: 4,
    NOT_FOUND: 5,
    UNAUTHORIZED: 6,
    SERVER_BUSY: 7,
    DISABLED: 8,
    INTERNAL_ERROR: 9
};

module.exports = {
    // TODO: method to convert code to string
    ...ERROR_CODE
};