package main;
 
import (
    "net"
    "time"
    "log"
    "strings"
)
 
func chkError(err error) {
    if err != nil {
        log.Fatal(err);
    }
}
 
//单独处理客户端的请求
func clientHandle(conn net.Conn) {
    //设置当客户端3分钟内无数据请求时，自动关闭conn
    conn.SetReadDeadline(time.Now().Add(time.Minute * 3));
    defer conn.Close();
 
    //循环的处理客户的请求
    for {
        data := make([]byte, 256);
        //从conn中读取数据
        n, err := conn.Read(data);
        //如果读取数据大小为0或出错则退出
        if n == 0 || err != nil {
            break;
        }
        //去掉两端空白字符
        cmd := strings.TrimSpace(string(data[0:n]));
        //发送给客户端的数据
        rep := "";
        if(cmd == "string") {
            rep = "hello,client \r\n";
        } else if (cmd == "time") {
            rep = time.Now().Format("2006-01-02 15:04:05");
        }
        //发送数据
        conn.Write([]byte(rep));
    }
}
 
func main() {
    tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
    chkError(err);
    tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
    chkError(err2);
    for {
        conn, err3 := tcplisten.Accept();
        if err3 != nil {
            continue;
        }
        go clientHandle(conn);
    }
}