use std::error::Error;
use std::fs::File;
use std::io::prelude::*;

pub mod client;

pub mod networkt;

pub struct Config {
    pub query: String,
    pub filename: String,
}

impl Config {
    fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() < 3 {
            return Err("not enough arguments");
        }

        let query = args[1].clone();
        let filename = args[2].clone();

        Ok(Config { query, filename })
    }
}

pub fn run(config: Config) -> Result<(), Box<Error>> {
    Ok(())
}

#[cfg(test)]
mod tests;
