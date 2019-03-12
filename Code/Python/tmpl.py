#!/usr/bin/env python3

import json
import logging
import os
import sys

import dotenv
from dotenv import find_dotenv, load_dotenv
from pydash import get, pick, uniq

logger = logging.getLogger('root-log')
logging.basicConfig(level=logging.INFO)

load_dotenv(find_dotenv())

test_env = os.getenv('TEST_ENV', 'false')

if __name__ == "__main__":
    logger.info(test_env)

    logger.debug('debug')
    logger.info('info')
    logger.warning('warn')
    logger.error('error')
