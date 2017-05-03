<?php
	
	/**
	* 配置：
		先建自己用户表，存储用户的etp钱包地址；建交易记录表，存储区块链上的交易记录；配置数据库链接，在init方法当中；配置区块链钱包的用户名，密码，在queryBlockTrades方法当中；
	
	* 开发流程：
		先创建钱包账户:createAccount（方法）;
		其余：
		    查出本地用户表中的所有etp地址，查出本地所有交易记录，迭代etp地址结果集，根据地址查处区块链上所有的交易记录，与本地交易记录对比，筛选出新的交易记录，保存到本地交易记录的数据库表中；listtxs（方法）	
		    发送交易：send（方法） 
	*/
	error_reporting(E_ERROR | E_WARNING | E_PARSE);	  
	
	/**
	*@date 2107-05-02
	*@description 查询钱包账户余额
	*@prama $username:客户的用户名，$password：客户设置的密码
	*/					
	function getbalance($username,$password){
		//请求参数（method:listtx,username:为钱包的用户名，password：为钱包的密码）  
		$post_data = array(  
		  "method" => "getbalance",  
		  "params" => [$username,$password]  
		);
		return _request($post_data);
	};
	
	/**
	*@date 2107-05-02
	*@description 发送交易
	*@prama $username:客户的用户名，$password：客户设置的密码，address：接收地址，quantity：金额
	*/					
	function send($username,$password,$address,$quantity){
		//请求参数（method:listtx,username:为钱包的用户名，password：为钱包的密码，address：接收地址，quantity：金额）  
		$post_data = array(  
		  "method" => "send",  
		  "params" => [$username,$password,$address,$quantity]  
		);
		return _request($post_data);
	};
	
	/**
	*@date 2107-05-02
	*@description 创建钱包账户
	*@prama $username:客户的用户名，$password：客户设置的密码
	*/					
	function createAccount($username,$password){
		//请求参数（method:listtx,username:为钱包的用户名，password：为钱包的密码）  
		$post_data = array(  
		  "method" => "getnewaccount",  
		  "params" => [$username,$password]  
		);
		return _request($post_data);
	};
	
	/**
	*@date 2107-05-02
	*@description 查询钱包里的所有交易记录（即区块链上的所有交易记录）;
	*/
	function queryBlockTrades($address){
		//请求参数（method:listtx,username:为钱包的用户名，password：为钱包的密码，address:为钱包的地址）  
		$post_data = array(  
		  "method" => "listtxs",  
		  "params" => ["username","password",$address.""]  
		);
		return _request($post_data);
	};
	
	/**
	*@date 2107-05-02
	*@description 处理交易相关
	*/
	function listtxs($conn){
		//充值地址查询结果（user表，本地自己建的数据库用户表）
		$addressSql = 'select id from fanwe_user'; 
		$addresses = select($addressSql,$conn);
		
		//本地数据库交易记录查询（trades表，本地自己建的数据库交易记录表）
		$tradeSql = 'select hash,address from trades'; 
		$trades = select($tradeSql,$conn);
		
		//迭代所有用户的地址
		while($address = mysql_fetch_array($addresses)) {
			//查询区块链上的交易记录
			$blockTrades = queryBlockTrades($address);
			$transactions = $blockTrades['transactions'];
			//迭代区块链上所有的交易记录
			while($blockTraderow = mysql_fetch_array($transactions)) {
				//迭代本地存储的交易记录
				while($localTraderow = mysql_fetch_array($trades)) {
					//判断本地与区块链上的交易hash是否相同
					if($blockTraderow['hash']!=$localTraderow['hash']){
						$outputs = $blockTraderow['hash'];
						//迭代区块链上的输出
						while($blockOutputrow = mysql_fetch_array($outputs)) {
							//判断输出与本地的地址是否相同
							if($localTraderow['address']!=$blockOutputrow['address']){
								//当满足交易hash与输出地址均不相同时,标示为新的交易,将新的交易插入本地交易数据库
								$localInsertSql = "insert into table(hash,address) values(".$blockTraderow['hash'].",".$blockOutputrow['address'].");";
								insert($localInsertSql,$conn);
							}
						}
					}
				}	
			}
		}
	};
	
	/**
	*@date 2107-05-02
	*@description 查询
	*@prama $sql：查询语句
	*/					
	function select($sql,$conn){
		//执行sql查询
		$result= mysql_query($sql, $conn);
		return $result;
	};
	
	/**
	*@date 2107-05-02
	*@description 插入数据库
	*/
	function insert($sql, $conn){
		//执行sql语句
		$result = mysql_query($sql, $conn);
		return $result;
	};
	
	/** 
	* 发送post请求 
	* @param string $url 请求地址 
	* @param params键值对数据 
	* @return string 
	*/ 
	function _request($params = array()){
		$params = json_encode($params);
		static $ch = null;
		if(is_null($ch)){
			$ch = curl_init();
			curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
			curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; BtcTrade PHP client; ' . php_uname('s') . '; PHP/' . phpversion() . ')');
		}
		# 参数设定
		curl_setopt($ch, CURLOPT_POSTFIELDS, $params);
		curl_setopt($ch, CURLOPT_URL, "http://localhost:8820/rpc");
		curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
		
		return curl_exec($ch);
	}
	
	
	/**
	*@date 2107-05-02
	*@description 初始化方法
	*/
	function init(){
		//数据库名称
		$mysql_name= 'fanwe';
		//数据库用户名
		$mysql_username = 'root';
		//数据库连接密码
		$mysql_password = 'root';
		//数据库ip地址
		$mysql_server_ip ="localhost";
		//连接到数据库
		$conn=mysql_connect($mysql_server_ip, $mysql_username,$mysql_password);
		//打开数据库
		mysql_select_db($mysql_name); 
		
		//创建钱包账户
		$account = createAccount("test32","test23");
		//$addressInsertSql = "insert into user(mnemonic,address) values(".$account['mnemonic'].",".$blockOutputrow['default-address'].");";
		//将创建出来的钱包地址保存到数据库当中
		//insert($addressInsertSql,$conn);
		
		//获得账户余额
		//send("test32","test23",'dfdbfg',500);
		
		//处理交易相关
		//listtxs($conn);
		
	};
	init();
	
?>