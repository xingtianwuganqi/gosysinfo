package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	containerName = "mysql"                    // 容器名称（不是镜像名）           // Docker MySQL 容器名
	dbUser        = "test_petrescue"           // MySQL 用户名
	dbPassword    = "pGfehdRfkin6X8pt"         // MySQL 密码
	dbName        = "test_petrescue"           // 要备份的数据库名称
	backupDir     = "/home/ubuntu/www/wwwroot" // 备份存储目录
	emailSender   = "your-email@example.com"   // 发件人邮箱
	emailPassword = "your-email-password"      // SMTP 授权码
	emailHost     = "smtp.example.com"         // SMTP 服务器
	emailPort     = 587                        // SMTP 端口
	emailReceiver = "receiver@example.com"     // 收件人邮箱
	GitRepoDir    = "/path/to/github/repo"     // GitHub 仓库路径
	GitToken      = "your_github_token"        // GitHub Token
	GitUser       = "your_github_username"
	GitRepoURL    = "https://github.com/your_github_username/your_repo.git"
)

func backupDatabase() error {
	// 确保目录存在
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		fmt.Println("❌ 创建备份目录失败:", err)
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("%s/mysql_backup_%s.sql", backupDir, timestamp)

	cmd := exec.Command("docker", "exec", containerName, "mysqldump",
		"-u"+dbUser, "-p"+dbPassword, dbName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ mysqldump 执行失败:", string(output))
		return err
	}

	if err := os.WriteFile(backupFile, output, 0644); err != nil {
		fmt.Println("❌ 写入备份文件失败:", err)
		return err
	}

	fmt.Println("✅ 数据库备份成功:", backupFile)
	return nil
}

func main() {
	if err := backupDatabase(); err != nil {
		fmt.Println("❌ 备份失败:", err)
	} else {
		fmt.Println("✅ 备份成功")
	}
}
