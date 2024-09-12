package utils

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func PromptPassword(prompt string) string {
	prompt1 := promptui.Prompt{
		Label:     prompt,
		Mask:      '*',
		IsConfirm: false,
	}
	result, err := prompt1.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return ""
	}
	return result
}
