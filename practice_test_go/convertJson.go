package main

import "time"

// Bat buoc de fiels in hoa vi can chuyen struct sang json thi build0in go can doc duoc de convert
type TodoItem struct {
    Id          int        `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Status      string     `json:"status"`
    CreatedAt   *time.Time `json:"created_at"` // de pointer vi neu DB ko co value se la nil
    UpdatedAt   *time.Time `json:"updated_at"` // neu ko de la pointer thi DB nil thi tren no van se ra gia tri
}

func main() {
    now := time.Now.Utc() // phai alocate vung nho thi moi lay con tro gan ben duoi duoc

    item := TodoItem {
        Id: 1,
        Title: "Test",
        Description: "This is test",
        Status: "Doing",
        CreatedAt: &now,
        UpdatedAt: nil,
    }

    data, err := json.Marshall(item)

    if err != nil {
        log.Fatalln(err)
    }

    log.Println(data)
}
