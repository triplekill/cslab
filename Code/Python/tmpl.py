#!/usr/bin/env python3

import argparse
import json
import logging
import os
import sys
from datetime import datetime

import dotenv
from dotenv import find_dotenv, load_dotenv
from pydash import get, pick, uniq

args = None
logger = logging.getLogger('root-log')
logging.basicConfig(level=logging.DEBUG)

load_dotenv(find_dotenv())

test_env = os.getenv('TEST_ENV', 'false')

def dateconv_json(o):
    if isinstance(o, datetime):
        return o.__str__()


def config():
    global args
    parser = argparse.ArgumentParser()
    parser.add_argument("--run", action="store_true",
                        help="actually run the update to database",)
    args = parser.parse_args()


if __name__ == "__main__":
    config()
    logger.debug(args.run)
    logger.info(test_env)

    logger.debug('debug')
    logger.info('info')
    logger.warning('warn')
    logger.error('error')
    query = {'k': datetime.now()}
    logger.debug(json.dumps(query, ensure_ascii=False, default=dateconv_json))
