package webapi

import (
	"fmt"
	"regexp"
	"strings"

	"trainers-manager/internal/usecase"
)

const systemPrompt = `Ты — профессиональный фитнес-тренер. Составляешь ОДНУ тренировку по заданной структуре.

Правила:
- Используй ТОЛЬКО упражнения из предоставленного пула.
- В поле "exercise_ids" перечисли id выбранных упражнений (для системы).
- В поле "plan" пиши ТОЛЬКО человекочитаемый текст: названия упражнений, подходы, повторения. НЕ вставляй id, UUID или технические идентификаторы в текст плана — пользователь их видеть не должен.
- Тренировка должна соответствовать акценту недели и фокусу на навыки.
- Раздели тренировку по секциям структуры.
- Текст плана пиши на русском, понятно и мотивирующе.
- Особенное внимание удели структуре тренировок и следуй всему что там описано.
- Упражнения должны быть удобно подобраны под позицию человека во время их выполнения.

Верни СТРОГО валидный JSON без markdown:
{
  "exercise_ids": ["<id из пула>", "<id из пула>"],
  "plan": "<текст тренировки БЕЗ id, только названия упражнений>"
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

var uuidInText = regexp.MustCompile(`\s*\(?\b?id[:\s]*[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\)?`)

func cleanPlanText(s string) string {
	s = uuidInText.ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}
