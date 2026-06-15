BASE="http://localhost:9045/v1/training"

# structure
curl -s -X POST "$BASE/structure" \
  -H "Content-Type: application/json" \
  -d '{"structure":"Кардио разминка, силовая часть, работа на выносливость, упражнения на пресс"}'

# exercise
curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"спина","description":"Подтягивания широким хватом, 4 подхода по 8-10 повторений", "position":"Стоя"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"спина","description":"Тяга штанги в наклоне, 4x10", "position":"Стоя"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"бицепс","description":"Сгибания на бицепс с гантелями, 3x12", "position":"Сидя"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"грудь","description":"Жим штанги лёжа, 4x8", "position":"Лежа"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"трицепс","description":"Отжимания на брусьях, 3x12", "position":"Стоя"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"пресс","description":"Скручивания на римском стуле, 3x15", "position":"Лежа"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"кардио","description":"Берпи, 4 раунда по 30 секунд", "position":"Стоя"}'

curl -s -X POST "$BASE/exercise" -H "Content-Type: application/json" \
  -d '{"muscle":"плечи","description":"Жим гантелей сидя, 4x10", "position":"Сидя"}'

# group
curl -s -X POST "$BASE/group" -H "Content-Type: application/json" \
  -d '{
    "name":"upper body",
    "accent_cycle":["спина, бицепс","грудь, трицепс","плечи, предплечья"],
    "skill_cycle":["баланс, взрывная сила","координация, выносливость"]
  }'
