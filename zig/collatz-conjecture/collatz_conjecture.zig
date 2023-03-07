pub const ComputationError = error{
    IllegalArgument,
};

pub fn steps(n: usize) anyerror!usize {
    if (n == 0) {
        return ComputationError.IllegalArgument;
    } else if (n == 1) {
        return 0;
    }
    return 1 + try (if (n & 1 == 1) steps(3 * n + 1) else steps(n / 2));
}
