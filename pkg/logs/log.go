package logs

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex // Log yazma işlemlerinde eşzamanlı erişimi kontrol etmek için mutex
var lastID int    // Son işlem ID'si

// Log yazma fonksiyonu
func WriteLog(action, username, status string) {
	mu.Lock()
	defer mu.Unlock()

	lastID++ // Her işlemde ID'yi artırıyoruz
	file, err := os.OpenFile("pkg/logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log dosyası açılırken hata oluştu:", err)
		return
	}
	defer file.Close()

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%d] [%s] - %s - Kullanıcı: %s, Durum: %s\n", lastID, currentTime, action, username, status)

	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Println("Log yazarken hata oluştu:", err)
	}
}
