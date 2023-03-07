import strutils as str 
import sequtils as sequ

proc hey*(s: string): string =
    var
        isUpper: bool
        isQuestion: bool
        stripped: string
        hasAlpha: bool
    stripped = s.strip()
    if str.isEmptyOrWhitespace(s):
        return "Fine. Be that way!"
    hasAlpha = stripped.any(str.isAlphaAscii)
    isUpper = hasAlpha and stripped.all(proc (c: char): bool =
        if str.isAlphaAscii(c):
            return str.isUpperAscii(c)
        return true
    )
    isQuestion = stripped.endsWith("?")
    if isUpper and isQuestion:
        return "Calm down, I know what I'm doing!"
    elif isUpper:
        return "Whoa, chill out!"
    elif isQuestion:
        return "Sure."
    else:
        return "Whatever."