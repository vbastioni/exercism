use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

fn sort(word: &str) -> (String, Vec<String>) {
    unsafe {
        let f = word.to_ascii_lowercase();
        let c = (&f).clone();
        let mut gs = c.graphemes(true).map(|s| s.to_owned()).collect::<Vec<_>>();
        gs.sort_unstable();
        (f, gs)
    }
}

pub fn anagrams_for<'a>(oword: &str, open: &'a [&'a str]) -> HashSet<&'a str> {
    let (lower, word) = sort(oword);
    HashSet::from_iter(
        open.iter()
            .map(|s| (*s, sort(*s)))
            .filter_map(|(s, (l, v))|  {
                if v == word && lower != l {
                    Some(s)
                } else {
                    None
                }
            }),
    )
}
