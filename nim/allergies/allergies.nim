type
    Allergy* = enum Eggs, Peanuts, Shellfish, Strawberries, Tomatoes, Chocolate, Pollen, Cats
    Allergens = set[Allergy]

proc allergies*(score: int): set[Allergy] = cast[Allergens](score)
proc isAllergicTo*(score: int, allergy: Allergy): bool = cast[bool](score and (1 shl ord(allergy)))
