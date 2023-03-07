use std::collections::HashMap;
use std::ops::AddAssign;
use std::sync::mpsc;
use std::sync::{Arc, Mutex};

use crossbeam::queue;

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let q = queue::ArrayQueue::new(input.len());
    for i in input {
        q.push(i.to_string()).expect("could not push to ArrayQueue");
    }
    let arc_q = Arc::new(q);

    let (send, recv) = mpsc::channel();
    // let 
    for _ in 0..worker_count {
        let sender = send.clone();
        let q = Arc::clone(&arc_q);
        std::thread::spawn(move || loop {
            match q.pop() {
                None => break,
                Some(s) => {
                    let mut h: HashMap<char, usize> =
                        s.chars()
                            .fold(HashMap::new(), |mut acc, cur| {
                                acc.entry(cur.to_ascii_lowercase())
                                    .or_insert(0)
                                    .add_assign(1);
                                acc
                            });
                    sender.send(h).expect("could not send, error here");
                }
            }
        });
    }

    HashMap::new()
}

pub fn frequency2(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    let id = Arc::new(Mutex::new((0, HashMap::new())));
    let v: Vec<_> = input.iter().map(|s| s.to_string()).collect();
    let inputs = Arc::new(v);
    let mut handles = vec![];
    for _ in 0..worker_count {
        let cloned = Arc::clone(&inputs);
        let cid = Arc::clone(&id);
        handles.push(std::thread::spawn(move || loop {
            let mut n = cid.lock().unwrap();
            if n.0 >= cloned.len() {
                break;
            }
            for c in (&cloned[n.0]).chars().filter(|c| c.is_alphabetic()) {
                n.1.entry(c.to_ascii_lowercase()).or_insert(0).add_assign(1);
            }
            n.0.add_assign(1);
        }));
    }
    for jh in handles.into_iter() {
        jh.join().unwrap();
    }
    let m = id.lock().unwrap();
    m.1.clone()
}
