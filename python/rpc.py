import httplib
import json
from functools import wraps
import socket
import inspect


def do_request(account_required):
    def do_request_(func):
        @wraps(func)
        def wrapper(self, *args, **kwargs):
            assert(isinstance(account_required, bool), 'account_required must be bool type')
            f = func(self, *args, **kwargs)
            params = list(args)
            if account_required:
                params = [self.account, self.passwd] + params
            infos = inspect.getargspec(func)
            options = {}
            reverse_args = infos.args[:]
            reverse_args.reverse()
            if infos.defaults:
                for idx, v in enumerate(reversed(infos.defaults)):
                    if v is None:
                        continue
                    params.append('--%s=%s' % (reverse_args[idx], v) )
            # params.extend(['--%s=%s' % (k, v) for k, v in kwargs.iteritems() if v is not None])
            method = func.__name__

            method = method.replace('_', '-')
            res = self.make_request(method=method, params = params)
            try:
                error_res = json.loads(res)
            except Exception as e:
                return f(res)
            if isinstance(error_res, dict) and error_res.get('error'):
                raise RuntimeError(error_res['error'])
            return f(res)
        return wrapper
    return do_request_


class Rpc:
    def __init__(self, account, passwd, host, port = 8820, debug =True):
        self.account = account
        self.passwd = passwd
        self.__host = host
        self.__port = port
        self.__method = None
        self.__debug = debug
        self.__avilable_methods_interface = {}
        self.conn = self.__create_connection()

    def __create_connection(self):
        conn = httplib.HTTPConnection(self.__host, self.__port)
        return conn

    @do_request(False)
    def fetch_height(self):
        return lambda x:x

    @do_request(False)
    def getblock(self, block_hash, is_json):
        return json.loads

    @do_request(False)
    def getbestblockhash(self):
        return str

    @do_request(True)
    def listbalances(self):
        return json.loads

    @do_request(True)
    def listtxs(self, address=None, height=None):
        return json.loads

    @do_request(True)
    def getbalance(self):
        return json.loads

    @do_request(False)
    def fetch_tx(self, hash_):
        return json.loads

    @do_request(False)
    def getnewaccount(self, account, passwd):
        return json.loads

    @do_request(True)
    def getnewaddress(self):
        return lambda x:x

    @do_request(False)
    def fetch_header(self, hash=None, height=100):
        return json.loads

    @do_request(True)
    def send(self, address, amount):
        return json.loads

    @do_request(True)
    def sendfrom(self, from_address, to_address, amount):
        return json.loads

    def make_request(self, method, params = []):
        max_try = 0
        while True:
            try:
                if not self.conn.sock:
                    max_try += 1
                    self.conn.connect()
            except Exception as e:
                if max_try > 10:
                    raise RuntimeError('try %s,%s' % (max_try, e))
                else:
                    print('reconnect to matching server(%s,%s) failed,%s' % (self.__host, self.__port, e))
            try:
                data = {'method':method, 'params':params}
                self.conn.request('POST', '/rpc', json.dumps(data))
                resp = self.conn.getresponse()
                txt = resp.read()
                return txt
            except httplib.error as e:
                self.conn.close()
                self.conn.sock = None
                continue
            except socket.error as e:
                self.conn.close()
                self.conn.sock = None
                continue


if __name__ == '__main__':
    r = Rpc('jiang1', 'jiang1', '172.16.52.128')
    balance = r.getbalance()
    print(balance)
    best_block_hash = r.getbestblockhash()
    print(best_block_hash)
    # block = r.getblock(best_block_hash + 'aa', True)
    block = r.getblock(best_block_hash, True)
    print(block)
    best_block_height = r.fetch_height()
    print(best_block_height)
    header = r.fetch_header()
    print(header)
    balances = r.listbalances()
    print(balances)
    # address_ = r.getnewaddress()
    # print(address_)
    txs = r.listtxs(address='MFpxQy4Y9soaUcTbPSHJxYbsThVCpR4WqM')
    print(txs)
    tx = r.fetch_tx('0347cbf42517b8d192372ae714dd016f46e29d20203129096e8427ef8c024afd')
    print(tx)
    tx = r.send('MDp7KV6xAia2ytg3V96RJjQzu9GPMyjsfi', 10000000000)
    print(tx)
