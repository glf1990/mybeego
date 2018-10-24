package db_mysql

import (
	"bufio"
	"bytes"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
)
var db *sql.DB

//注意方法名大写，就是public
func Create_db(ip string,port string,dbName string,userName string,password string) (error) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	var err error
	db, err = sql.Open("mysql", path)
	if err != nil {
		return err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	/*if err := db.Ping(); err != nil{
		//fmt.Println("opon database fail")
		return err
	}*/
	return nil
}

func Create_db_by_ssh(database_ip string,database_port int,dbName string,userName string,
	password string,database_LocalIP string,database_LocalPort int,sshServerHost string,sshServerPort int,sshUserName string,
	sshPrivateKeyFilePath string,sshKeyPassphrase string)(error){
	/*const(
		// ssh connection related data
		//sshServerHost     = "203.116.243.3" //跳板机IP或者域名
		//sshServerPort     = 22 //跳板机ssh的端口好
		//sshUserName       = "ld-xuefeng_song" //登录跳板机的用户名
		//sshPrivateKeyFile = "/Users/xuefeng.song/.ssh/id_rsa" //私钥的全路径 exported as OpenSSH key from .ppk
		//sshKeyPassphrase  = "" // key file encrytion password //私钥的密码

		// ssh tunneling related data
		//database_LocalIP  = "127.0.0.1" // local localhost ip (client side) 映射本机的IP地址一般用127.0.0.1
		//sshLocalPort  = 3308        // local port used to forward the connection 映射本机的端口
		sshRemoteHost = "10.65.136.147" // remote local ip (server side) 数据库的IP服务器地址
		sshRemotePort = 3306        // remote MySQL port 数据库的端口

		// MySQL access data
		mySqlUsername = "root" //登录数据库的用户名
		mySqlPassword = "" //登录数据库的密码
		mySqlDatabase = "" //登录的数据库
	)*/
	tunnel ,errr := sshTunnel(database_LocalIP ,database_LocalPort ,sshServerHost ,sshServerPort ,database_ip,
		database_port ,sshUserName ,sshPrivateKeyFilePath ,sshKeyPassphrase )  // Initialize sshTunnel
	go tunnel.Start()     // Start the sshTunnel
	if errr != nil {
		return errr
	}
	// Declare the dsn (aka database connection string)
	// dsn := "opensim:h0tgrits@tcp(localhost:9000)/opensimdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		userName, password, database_LocalIP, database_LocalPort, dbName)

	// Open the database
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	return nil
}
func Close_db(){
	db.Close()
}
func Select(sql string)(*sql.Rows,error) {
	rows, err := db.Query(sql)
	if err != nil {
		//println(err)
		return nil,err
	}
	return rows,nil
}
func Update_getAffectedRowNum(sql string)(int64,error){
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		//println(err)
		return 0,err
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		//println(err)
		return 0,err
	}
	//res, err := stmt.Exec(params)
	res, err := stmt.Exec()
	if err != nil {
		//println(err)
		return 0,err
	}
	tx.Commit()
	//id, err := res.LastInsertId()
	affect, err := res.RowsAffected()
	//fmt.Println(affect)
	if err != nil {
		//println(err)
		return 0,err
	}
	return affect,nil
}
func Insert_getLastId(sql string)(int64,error){
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		//println(err)
		return 0,err
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		//println(err)
		return 0,err
	}
	//res, err := stmt.Exec(params)
	res, err := stmt.Exec()
	if err != nil {
		//println(err)
		return 0,err
	}
	tx.Commit()
	id, err := res.LastInsertId()
	//affect, err := res.RowsAffected()
	//fmt.Println(affect)
	if err != nil {
		//println(err)
		return 0,err
	}
	return id,nil
}
func Get_json(rows *sql.Rows) (string, error) {
	//defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
			}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
				} else {
					v = val
					}
			entry[col] = v
			}
		tableData = append(tableData, entry)
		}
	jsonData, err := json.Marshal(tableData)
	if err != nil {   return "", err    }
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}
/* 引用示例
err := db_mysql.Create_db(ip,port,dbName,userName,password)
defer db_mysql.Close_db()
if err!=nil{
fmt.Println(err)
}
rows,err:=db_mysql.Get_data("select * from test")
defer rows.Close()
if err!=nil{
fmt.Println(err)
}
fmt.Println("hello world!!!")
if !rows.Next() {
        return
    }
for rows.Next() {
var id int
var name string
err = rows.Scan(&id,&name)
if err != nil {
panic(err)
}
fmt.Println(id,name)
}

*/
// Simple mySql error handling (yet to implement)
//func dbErrorHandler(err error) {
//	switch err := err.(type) {
//	default:
//		fmt.Printf("Error %s\n", err)
//		os.Exit(-1)
//	}
//}
// Define an endpoint with ip and port
type Endpoint struct {
	Host string
	Port int
}

// Returns an endpoint as ip:port formatted string
func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

// Define the endpoints along the tunnel
type SSHtunnel struct {
	Local  *Endpoint
	Server *Endpoint
	Remote *Endpoint
	Config *ssh.ClientConfig
}

// Start the tunnel
func (tunnel *SSHtunnel) Start() error {
	listener, err := net.Listen("tcp", tunnel.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go tunnel.forward(conn)
	}
}

// Port forwarding
func (tunnel *SSHtunnel) forward(localConn net.Conn) {
	// Establish connection to the intermediate server
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	// access the target server
	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	// Transfer the data between  and the remote server
	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

// Decrypt encrypted PEM key data with a passphrase and embed it to key prefix
// and postfix header data to make it valid for further private key parsing.
func DecryptPEMkey(buffer []byte, passphrase string) []byte {
	block, _ := pem.Decode(buffer)
	der, err := x509.DecryptPEMBlock(block, []byte(passphrase))
	if err != nil {
		fmt.Println("decrypt failed: ", err)
	}
	encoded := base64.StdEncoding.EncodeToString(der)
	encoded = "-----BEGIN RSA PRIVATE KEY-----\n" + encoded +
		"\n-----END RSA PRIVATE KEY-----\n"
	return []byte(encoded)
}

// Get the signers from the OpenSSH key file (.pem) and return them for use in
// the Authentication method. Decrypt encrypted key data with the passphrase.
func PublicKeyFile(file string, passphrase string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	if bytes.Contains(buffer, []byte("ENCRYPTED")) {
		// Decrypt the key with the passphrase if it has been encrypted
		buffer = DecryptPEMkey(buffer, passphrase)
	}

	// Get the signers from the key
	signers, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(signers)
}

// Define the ssh tunnel using its endpoint and config data
func sshTunnel(sshLocalHost string,sshLocalPort int,sshServerHost string,sshServerPort int,sshRemoteHost string,sshRemotePort int,sshUserName string,sshPrivateKeyFilePath string,sshKeyPassphrase string) (*SSHtunnel,error) {
	localEndpoint := &Endpoint{
		Host: sshLocalHost,
		Port: sshLocalPort,
	}

	serverEndpoint := &Endpoint{
		Host: sshServerHost,
		Port: sshServerPort,
	}

	remoteEndpoint := &Endpoint{
		Host: sshRemoteHost,
		Port: sshRemotePort,
	}
	var hostKey ssh.PublicKey
	hostKey, err := getHostKey(sshServerHost)
	//fmt.Println(hostKey)
	if err != nil {
		return nil,err
	}
	sshConfig := &ssh.ClientConfig{
		User: sshUserName,
		Auth: []ssh.AuthMethod{
			PublicKeyFile(sshPrivateKeyFilePath, sshKeyPassphrase)},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	return &SSHtunnel{
		Config: sshConfig,
		Local:  localEndpoint,
		Server: serverEndpoint,
		Remote: remoteEndpoint,
	},nil
}
func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}
	return hostKey, nil
}