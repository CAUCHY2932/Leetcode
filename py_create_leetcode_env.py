# coding:utf-8

import os
import sys
import logging


# get current file path

# dirname, filename = os.path.split(os.path.abspath(sys.argv[0]))
# os.path.realpath(sys.argv[0]) # current file path

dirname, _ = os.path.split(os.path.abspath(sys.argv[0]))

# config area

LOG_FILE = "leetcode.log"
LOG_FORMAT = "%(asctime)s - %(levelname)s - %(message)s"
DATE_FORMAT = "%Y-%m-%d %H:%M:%S %p"
FILE_HEADER = "package leetcode{}\n"


logging.basicConfig(level=logging.DEBUG,
                    filename=LOG_FILE,
                    format=LOG_FORMAT,
                    datefmt=DATE_FORMAT,
                    )


def main():
    num = input_and_analyse()
    create_code_env(num=num)
    logging.debug("have create code env")
    print("have create code env successfully!")


def input_and_analyse():
    num = input("please input the leetcode num:\n") # python2和python3不同
    if type(num) == type(0):
        print("num")
        return str(num).zfill(4)
    else:
        num = num.strip()
        if num == "":
            print("invalid, num")
            logging.debug("invalid num")
            exit
        else:
            logging.debug("have input num: {}".format(num))
            return num.zfill(4)


def create_code_env(num):
    print("-"*10+"have input ", num)
    logging.debug("is create code env...")
    logging.debug("change dir of %s..."%dirname)
    os.chdir(dirname)
    mkdir(folder_name=os.path.join("%s/src"%dirname, "leetcode{}".format(num)))
    touch_file(filename="./src/leetcode{}/leetcode{}.go".format(num, num),
               content=FILE_HEADER.format(num))
    touch_file(filename="./src/leetcode{}/leetcode{}_test.go".format(num, num),
               content=FILE_HEADER.format(num))
    pass


def mkdir(folder_name):
    try:
        os.mkdir(folder_name)
    except Exception as e:
        logging.debug("mkdir {} error, reason for {}".format(folder_name, e))
    else:
        logging.debug("have mkdir {}".format(folder_name))


def touch_file(filename, content):
    with open(file=filename, mode="w", encoding="utf-8")as f:
        f.write(content)
        logging.debug("have write content to {}".format(filename))


if __name__ == "__main__":
    main()
