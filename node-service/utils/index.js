const responseGenerator = (msg, success, data = []) => {
    return {
        data,
        msg,
        success,
    };
};

module.exports = {
    responseGenerator,
};
