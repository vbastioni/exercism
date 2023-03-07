pub fn squareOfSum(n: usize) usize {
    return (n * (1 + n) / 2) * (n * (1 + n) / 2);
}

pub fn sumOfSquares(n: usize) usize {
    return (n * (n + 1) * (2 * n + 1)) / 6;
}

pub fn differenceOfSquares(n: usize) usize {
    return squareOfSum(n) - sumOfSquares(n);
}
