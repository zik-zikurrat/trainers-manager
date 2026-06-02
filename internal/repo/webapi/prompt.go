package webapi

import (
	"fmt"
	"strings"

	"trainers-manager/internal/usecase"
)

const systemPrompt = `Ты — профессиональный фитнес-тренер. Составляешь ОДНУ тренировку по заданной структуре.

Правила:
- Используй ТОЛЬКО упражнения из предоставленного пула, ссылайся на них по точному id.
- НЕ придумывай упражнения и id, которых нет в пуле.
- Тренировка должна соответствовать акценту недели и фокусу на навыки.
- Раздели тренировку по секциям структуры.
- Текст плана пиши на русском, понятно и мотивирующе.

Верни СТРОГО валидный JSON без markdown, без пояснений, ровно такого вида:
{
  "exercise_ids": ["<id из пула>", "<id из пула>"],
  "plan": "<читаемый текст тренировки для пользователя>"
}`

func buildUserPrompt(p usecase.GeneratePrompt) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Структура тренировки: %s\n", p.Structure)
	fmt.Fprintf(&b, "Акцент недели: %s\n", p.Accent)
	fmt.Fprintf(&b, "Фокус на навыки: %s\n\n", p.Skills)

	if len(p.Recent) > 0 {
		b.WriteString("Недавние тренировки (НЕ повторяй их акцент и набор):\n")
		for _, r := range p.Recent {
			fmt.Fprintf(&b, "- акцент: %s | навыки: %s\n", r.Accent, r.Skills)
		}
		b.WriteString("\n")
	}

	b.WriteString("Пул доступных упражнений (выбирай ТОЛЬКО отсюда):\n")
	for _, e := range p.Pool {
		fmt.Fprintf(&b, "- id=%s | мышца: %s | %s\n", e.ID, e.Muscle, e.Description)
	}

	return b.String()
}
