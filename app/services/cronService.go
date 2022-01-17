package services

import (
	"go-api/global"
	"go-api/tool"

	cron "github.com/18211167516/robfig-cron/v3"
)

type Task struct {
	ID   cron.EntryID `json:"id"`
	Name string       `json:"name"`
	Spec string       `json:"spec"`
	Next string       `json:"next"`
	Prev string       `json:"prev"`
}

func GetCronList() tool.M {
	data := make(map[string]interface{})
	crons := global.CRON.Entries()

	if count := len(crons); count > 0 {
		var tasks []Task
		for _, v := range crons {
			if v.Valid() {
				task := Task{
					ID:   v.ID,
					Name: v.Name,
					Spec: v.Spec,
					Next: v.Next.Format("2006-01-02 15:04:05"),
					Prev: v.Prev.Format("2006-01-02 15:04:05"),
				}
				tasks = append(tasks, task)
			}
		}
		data["lists"] = tasks
		data["total"] = count
		return tool.DataReturn(true, "查询成功", data)
	}

	return tool.DataReturn(false, "暂无数据", nil)
}

func RunCron(id string) tool.M {
	Entryid := cron.EntryID(tool.StringToInt(id))

	entry := global.CRON.Entry(Entryid)
	if entry.Valid() {
		global.CRON.RunJob(&entry)
		return tool.DataReturn(true, "执行成功", nil)
	}

	return tool.DataReturn(false, "任务不存在", nil)
}

func DelCron(id string) tool.M {
	Entryid := cron.EntryID(tool.StringToInt(id))
	entry := global.CRON.Entry(Entryid)
	if entry.Valid() {
		global.CRON.Remove(Entryid)
		return tool.DataReturn(true, "删除成功", nil)
	}

	return tool.DataReturn(false, "任务不存在", nil)
}
