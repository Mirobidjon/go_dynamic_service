package http

var CustomErrorMessages = map[string]map[string]string{
	// DEFAULT ERROR MESSAGE
	"default_error_message": {
		"uz": "Kutilmagan xatolik",
		"oz": "Кутилмаган хатолик",
		"ru": "Неизвестная ошибка",
	},

	// BAD REQUEST
	"unauthorized": {
		"uz": "Avtorizatsiyadan o'tmagan",
		"oz": "Авторизатсиядан ўтмаган",
		"ru": "Не авторизован",
	},

	// BODY PARSE
	"body_parse": {
		"uz": "Body parse xatoligi",
		"oz": "Body parse хатолиги",
		"ru": "Ошибка парсинга body",
	},

	// QUERY PARSE
	"query_parse": {
		"uz": "Query parse xatoligi",
		"oz": "Query parse хатолиги",
		"ru": "Ошибка парсинга query",
	},

	// OFFSET LIMIT
	"get_offset": {
		"uz": "Offsetni olishda xatolik yuz berdi",
		"oz": "Offsetни олишда хатолик юз берди",
		"ru": "Ошибка получения offset",
	},
	"get_limit": {
		"uz": "Limitni olishda xatolik yuz berdi",
		"oz": "Limitни олишда хатолик юз берди",
		"ru": "Ошибка получения limit",
	},

	// CONVERT MAP TO STRUCT
	"convert_map_to_struct": {
		"uz": "Mapni structga o'girishda xatolik",
		"oz": "Mapни structга ўгиришда хатолик",
		"ru": "Ошибка конвертации map в struct",
	},

	// INVALID UUID
	"invalid_uuid": {
		"uz": "UUID noto'g'ri",
		"oz": "UUID нотўғри",
		"ru": "Неверный UUID",
	},

	// INVALID SLUG
	"invalid_slug": {
		"uz": "Slug noto'g'ri",
		"oz": "Slug нотўғри",
		"ru": "Неверный Slug",
	},

	// INVALID ID
	"invalid_id": {
		"uz": "ID noto'g'ri",
		"oz": "ID нотўғри",
		"ru": "Неверный ID",
	},

	// BYTE TO PROTO
	"byte_to_proto": {
		"uz": "Byte to proto xatoligi",
		"oz": "Byte to proto хатолиги",
		"ru": "Ошибка byte to proto",
	},

	// PROTO TO STRUCT
	"proto_to_struct": {
		"uz": "Proto to struct xatoligi",
		"oz": "Proto to struct хатолиги",
		"ru": "Ошибка proto to struct",
	},

	// JSON MARSHAL
	"json_marshal": {
		"uz": "Json marshal xatoligi",
		"oz": "Json marshal хатолиги",
		"ru": "Ошибка json marshal",
	},

	// 	INVALID TOKEN
	"invalid_token": {
		"uz": "Token noto'g'ri",
		"oz": "Token нотўғри",
		"ru": "Неверный Token",
	},

	// ENTITY
	"unmarshal_entity": {
		"uz": "Entityni unmarshal qilishda xatolik yuz berdi",
		"oz": "Entityни unmarshal қилишда хатолик юз берди",
		"ru": "Ошибка при unmarshal entity",
	},
	"create_entity": {
		"uz": "Entityni yaratishda xatolik yuz berdi",
		"oz": "Entityни яратишда хатолик юз берди",
		"ru": "Ошибка при создании entity",
	},
	"update_entity": {
		"uz": "Entityni yangilashda xatolik yuz berdi",
		"oz": "Entityни янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении entity",
	},
	"delete_entity": {
		"uz": "Entityni o'chirishda xatolik yuz berdi",
		"oz": "Entityни ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении entity",
	},
	"get_entity": {
		"uz": "Entityni olishda xatolik yuz berdi",
		"oz": "Entityни олишда хатолик юз берди",
		"ru": "Ошибка при получении entity",
	},
	"get_entity_list": {
		"uz": "Entitylar ro'yxatini olishda xatolik yuz berdi",
		"oz": "Entityлар рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка entity",
	},
	"join_entity": {
		"uz": "Entityni join qilishda xatolik yuz berdi",
		"oz": "Entityни join қилишда хатолик юз берди",
		"ru": "Ошибка при join entity",
	},

	// CASHBOX
	"create_cashbox": {
		"uz": "Kassani yaratishda xatolik yuz berdi",
		"oz": "Кассани яратишда хатолик юз берди",
		"ru": "Ошибка при создании кассы",
	},
	"get_cashbox": {
		"uz": "Kassani olishda xatolik yuz berdi",
		"oz": "Кассани олишда хатолик юз берди",
		"ru": "Ошибка при получении кассы",
	},
	"get_cashbox_list": {
		"uz": "Kassalar ro'yxatini olishda xatolik yuz berdi",
		"oz": "Кассалар рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка касс",
	},
	"update_cashbox": {
		"uz": "Kassani yangilashda xatolik yuz berdi",
		"oz": "Кассани янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении кассы",
	},
	"delete_cashbox": {
		"uz": "Kassani o'chirishda xatolik yuz berdi",
		"oz": "Кассани ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении кассы",
	},

	// CLIENT PLATFORM
	"create_client_platform": {
		"uz": "Platformani yaratishda xatolik yuz berdi",
		"oz": "Платформани яратишда хатолик юз берди",
		"ru": "Ошибка при создании платформы",
	},
	"get_client_platform": {
		"uz": "Platformani olishda xatolik yuz berdi",
		"oz": "Платформани олишда хатолик юз берди",
		"ru": "Ошибка при получении платформы",
	},
	"get_client_platform_list": {
		"uz": "Platformalar ro'yxatini olishda xatolik yuz berdi",
		"oz": "Платформалар рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка платформ",
	},
	"update_client_platform": {
		"uz": "Platformani yangilashda xatolik yuz berdi",
		"oz": "Платформани янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении платформы",
	},
	"delete_client_platform": {
		"uz": "Platformani o'chirishda xatolik yuz berdi",
		"oz": "Платформани ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении платформы",
	},

	// CLIENT TYPE
	"create_client_type": {
		"uz": "Klient tipini yaratishda xatolik yuz berdi",
		"oz": "Клиент типини яратишда хатолик юз берди",
		"ru": "Ошибка при создании типа клиента",
	},
	"get_client_type": {
		"uz": "Klient tipini olishda xatolik yuz berdi",
		"oz": "Клиент типини олишда хатолик юз берди",
		"ru": "Ошибка при получении типа клиента",
	},
	"get_client_type_list": {
		"uz": "Klient tipi ro'yxatini olishda xatolik yuz berdi",
		"oz": "Клиент типи рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка типов клиентов",
	},
	"update_client_type": {
		"uz": "Klient tipini yangilashda xatolik yuz berdi",
		"oz": "Клиент типини янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении типа клиента",
	},
	"delete_client_type": {
		"uz": "Klient tipini o'chirishda xatolik yuz berdi",
		"oz": "Клиент типини ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении типа клиента",
	},

	// CLIENT
	"add_client": {
		"uz": "Klient qo'shishda xatolik yuz berdi",
		"oz": "Клиент қўшишда хатолик юз берди",
		"ru": "Ошибка добавления клиента",
	},
	"get_client_matrix": {
		"uz": "Klient matrixini olishda xatolik yuz berdi",
		"oz": "Клиент матриксини олишда хатолик юз берди",
		"ru": "Ошибка получения матрицы клиента",
	},
	"update_client": {
		"uz": "Klientni yangilashda xatolik yuz berdi",
		"oz": "Клиентни янгилашда хатолик юз берди",
		"ru": "Ошибка обновления клиента",
	},
	"remove_client": {
		"uz": "Klientni o'chirishda xatolik yuz berdi",
		"oz": "Клиентни ўчиришда хатолик юз берди",
		"ru": "Ошибка удаления клиента",
	},

	// RELATION
	"add_relation": {
		"uz": "Relation qo'shishda xatolik yuz berdi",
		"oz": "Relation қўшишда хатолик юз берди",
		"ru": "Ошибка добавления relation",
	},
	"update_relation": {
		"uz": "Relation yangilashda xatolik yuz berdi",
		"oz": "Relation янгилашда хатолик юз берди",
		"ru": "Ошибка обновления relation",
	},
	"remove_relation": {
		"uz": "Relation o'chirishda xatolik yuz berdi",
		"oz": "Relation ўчиришда хатолик юз берди",
		"ru": "Ошибка удаления relation",
	},

	// USER INFO
	"add_user_info": {
		"uz": "User info qo'shishda xatolik yuz berdi",
		"oz": "User info қўшишда хатолик юз берди",
		"ru": "Ошибка добавления user info",
	},
	"update_user_info": {
		"uz": "User info yangilashda xatolik yuz berdi",
		"oz": "User info янгилашда хатолик юз берди",
		"ru": "Ошибка обновления user info",
	},
	"remove_user_info": {
		"uz": "User info o'chirishda xatolik yuz berdi",
		"oz": "User info ўчиришда хатолик юз берди",
		"ru": "Ошибка удаления user info",
	},

	// COMPANY
	"create_company": {
		"uz": "Kompaniyani yaratishda xatolik yuz berdi",
		"oz": "Компанияни яратишда хатолик юз берди",
		"ru": "Ошибка при создании компании",
	},
	"get_company": {
		"uz": "Kompaniyani olishda xatolik yuz berdi",
		"oz": "Компанияни олишда хатолик юз берди",
		"ru": "Ошибка при получении компании",
	},
	"get_company_list": {
		"uz": "Kompaniyalar ro'yxatini olishda xatolik yuz berdi",
		"oz": "Компаниялар рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка компаний",
	},
	"update_company": {
		"uz": "Kompaniyani yangilashda xatolik yuz berdi",
		"oz": "Компанияни янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении компании",
	},
	"delete_company": {
		"uz": "Kompaniyani o'chirishda xatolik yuz berdi",
		"oz": "Компанияни ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении компании",
	},

	// FM
	"create_fm": {
		"uz": "FM yaratishda xatolik yuz berdi",
		"oz": "FM яратишда хатолик юз берди",
		"ru": "Ошибка при создании FM",
	},
	"get_fm": {
		"uz": "FM olishda xatolik yuz berdi",
		"oz": "FM олишда хатолик юз берди",
		"ru": "Ошибка при получении FM",
	},
	"get_fm_list": {
		"uz": "FM ro'yxatini olishda xatolik yuz berdi",
		"oz": "FM рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка FM",
	},
	"update_fm": {
		"uz": "FM yangilashda xatolik yuz berdi",
		"oz": "FM янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении FM",
	},
	"delete_fm": {
		"uz": "FM o'chirishda xatolik yuz berdi",
		"oz": "FM ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении FM",
	},

	// GROUP
	"create_group": {
		"uz": "Group yaratishda xatolik yuz berdi",
		"oz": "Group яратишда хатолик юз берди",
		"ru": "Ошибка при создании Group",
	},
	"get_group": {
		"uz": "Group olishda xatolik yuz berdi",
		"oz": "Group олишда хатолик юз берди",
		"ru": "Ошибка при получении Group",
	},
	"get_group_list": {
		"uz": "Group ro'yxatini olishda xatolik yuz berdi",
		"oz": "Group рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка Group",
	},
	"update_group": {
		"uz": "Group yangilashda xatolik yuz berdi",
		"oz": "Group янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении Group",
	},
	"delete_group": {
		"uz": "Group o'chirishda xatolik yuz berdi",
		"oz": "Group ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Group",
	},
	"get_full_group": {
		"uz": "Full groupni olishda xatolik yuz berdi",
		"oz": "Full groupни олишда хатолик юз берди",
		"ru": "Ошибка при получении Full group",
	},

	// FIELD
	"create_field": {
		"uz": "Field yaratishda xatolik yuz berdi",
		"oz": "Field яратишда хатолик юз берди",
		"ru": "Ошибка при создании Field",
	},
	"get_field": {
		"uz": "Field olishda xatolik yuz berdi",
		"oz": "Field олишда хатолик юз берди",
		"ru": "Ошибка при получении Field",
	},
	"get_field_list": {
		"uz": "Field ro'yxatini olishda xatolik yuz berdi",
		"oz": "Field рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка Field",
	},
	"update_field": {
		"uz": "Field yangilashda xatolik yuz berdi",
		"oz": "Field янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении Field",
	},
	"delete_field": {
		"uz": "Field o'chirishda xatolik yuz berdi",
		"oz": "Field ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Field",
	},

	// ROLE
	"create_role": {
		"uz": "Role yaratishda xatolik yuz berdi",
		"oz": "Role яратишда хатолик юз берди",
		"ru": "Ошибка при создании Role",
	},
	"get_role": {
		"uz": "Role olishda xatolik yuz berdi",
		"oz": "Role олишда хатолик юз берди",
		"ru": "Ошибка при получении Role",
	},
	"get_role_list": {
		"uz": "Role ro'yxatini olishda xatolik yuz berdi",
		"oz": "Role рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка Role",
	},
	"update_role": {
		"uz": "Role yangilashda xatolik yuz berdi",
		"oz": "Role янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении Role",
	},
	"delete_role": {
		"uz": "Role o'chirishda xatolik yuz berdi",
		"oz": "Role ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Role",
	},

	// PERMISSION
	"create_permission": {
		"uz": "Permission yaratishda xatolik yuz berdi",
		"oz": "Permission яратишда хатолик юз берди",
		"ru": "Ошибка при создании Permission",
	},
	"get_permission": {
		"uz": "Permission olishda xatolik yuz berdi",
		"oz": "Permission олишда хатолик юз берди",
		"ru": "Ошибка при получении Permission",
	},
	"get_permission_list": {
		"uz": "Permission ro'yxatini olishda xatolik yuz berdi",
		"oz": "Permission рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка Permission",
	},
	"update_permission": {
		"uz": "Permission yangilashda xatolik yuz berdi",
		"oz": "Permission янгилашда хатолик юз берди",
		"ru": "Ошибка при обновлении Permission",
	},
	"delete_permission": {
		"uz": "Permission o'chirishda xatolik yuz berdi",
		"oz": "Permission ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Permission",
	},

	// SCOPE
	"create_user_scope": {
		"uz": "User scope yaratishda xatolik yuz berdi",
		"oz": "User scope яратишда хатолик юз берди",
		"ru": "Ошибка при создании User scope",
	},
	"get_scope_list": {
		"uz": "Scope ro'yxatini olishda xatolik yuz berdi",
		"oz": "Scope рўйхатини олишда хатолик юз берди",
		"ru": "Ошибка при получении списка Scope",
	},
	"add_permission_scope": {
		"uz": "Permission scope qo'shishda xatolik yuz berdi",
		"oz": "Permission scope қўшишда хатолик юз берди",
		"ru": "Ошибка при добавлении Permission scope",
	},
	"remove_permission_scope": {
		"uz": "Permission scope o'chirishda xatolik yuz berdi",
		"oz": "Permission scope ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Permission scope",
	},

	// ROLE PERMISSION
	"add_role_permission": {
		"uz": "Role permission qo'shishda xatolik yuz berdi",
		"oz": "Role permission қўшишда хатолик юз берди",
		"ru": "Ошибка при добавлении Role permission",
	},
	"add_role_permissions": {
		"uz": "Role permissions qo'shishda xatolik yuz berdi",
		"oz": "Role permissions қўшишда хатолик юз берди",
		"ru": "Ошибка при добавлении Role permissions",
	},
	"remove_role_permission": {
		"uz": "Role permission o'chirishda xatolik yuz berdi",
		"oz": "Role permission ўчиришда хатолик юз берди",
		"ru": "Ошибка при удалении Role permission",
	},

	// SESSION
	"login": {
		"uz": "Login qilishda xatolik yuz berdi",
		"oz": "Login қилишда хатолик юз берди",
		"ru": "Ошибка при логине",
	},
	"logout": {
		"uz": "Logout qilishda xatolik yuz berdi",
		"oz": "Logout қилишда хатолик юз берди",
		"ru": "Ошибка при логауте",
	},
	"refresh_token": {
		"uz": "Refresh token qilishda xatolik yuz berdi",
		"oz": "Refresh token қилишда хатолик юз берди",
		"ru": "Ошибка при обновлении токена",
	},
	"has_access": {
		"uz": "Has access xatolik yuz berdi",
		"oz": "Has access хатолик юз берди",
		"ru": "Ошибка при получении доступа",
	},

	// PRODUCT
	"add_product": {
		"uz": "Mahsulot qo'shishda xatolik",
		"oz": "Mаҳсулот қўшишда хатолик",
		"ru": "Ошибка добавления продукта",
	},
	"update_product": {
		"uz": "Mahsulotni yangilashda xatolik",
		"oz": "Mаҳсулотни янгилашда хатолик",
		"ru": "Ошибка обновления продукта",
	},
	"search_product": {
		"uz": "Mahsulotni qidirishda xatolik",
		"oz": "Mаҳсулотни қидиришда хатолик",
		"ru": "Ошибка поиска продукта",
	},
	"product_name_is_required": {
		"uz": "Mahsulot nomi to'ldirilmagan",
		"oz": "Маҳсулот номи тўлдирилмаган",
		"ru": "Название продукта не заполнено",
	},
	"mxik_code_is_required": {
		"uz": "Mahsulot MXIK kodi to'ldirilmagan",
		"oz": "Маҳсулот MXIK коди тўлдирилмаган",
		"ru": "MXIK код продукта не заполнен",
	},
	"barcode_is_required": {
		"uz": "Mahsulot shtrix kodi to'ldirilmagan",
		"oz": "Маҳсулот штрих коди тўлдирилмаган",
		"ru": "Штрих код продукта не заполнен",
	},
	"income_price_is_required": {
		"uz": "Mahsulot kiruvchi narxi to'ldirilmagan",
		"oz": "Маҳсулот кирувчи нархи тўлдирилмаган",
		"ru": "Цена продукта не заполнена",
	},
	"sell_price_is_required": {
		"uz": "Mahsulot sotuv narxi to'ldirilmagan",
		"oz": "Маҳсулот сотув нархи тўлдирилмаган",
		"ru": "Цена продукта не заполнена",
	},
	"remaining_count_is_required": {
		"uz": "Mahsulot qoldig'i to'ldirilmagan",
		"oz": "Маҳсулот қолдиғи тўлдирилмаган",
		"ru": "Остаток продукта не заполнен",
	},
	"vat_percent_is_required": {
		"uz": "Mahsulot QQS foizi to'ldirilmagan",
		"oz": "Маҳсулот ҚҚС фоизи тўлдирилмаган",
		"ru": "Процент НДС продукта не заполнен",
	},
	"minimum_unit_is_required": {
		"uz": "Mahsulot minimal birlik to'ldirilmagan",
		"oz": "Маҳсулот минимал бирлик тўлдирилмаган",
		"ru": "Минимальная единица продукта не заполнена",
	},
	"discount_type_is_required": {
		"uz": "Mahsulot chegirma turi to'ldirilmagan",
		"oz": "Маҳсулот чегирма тури тўлдирилмаган",
		"ru": "Тип скидки продукта не заполнен",
	},
	"discount_type_is_invalid": {
		"uz": "Mahsulot chegirma turi noto'g'ri",
		"oz": "Маҳсулот чегирма тури нотўғри",
		"ru": "Тип скидки продукта неверен",
	},
	"discount_value_is_invalid": {
		"uz": "Mahsulot chegirma qiymati noto'g'ri",
		"oz": "Маҳсулот чегирма қиймати нотўғри",
		"ru": "Значение скидки продукта неверен",
	},
	"fm_id_is_required": {
		"uz": "Fiscal modul bo'sh",
		"oz": "Фискал модуль бўш",
		"ru": "Фискальный модуль пуст",
	},
	"company_id_is_required": {
		"uz": "Kompaniya idsi to'ldirilmagan",
		"oz": "Компания идси тўлдирилмаган",
		"ru": "Ид компании не заполнен",
	},
	"cashbox_id_is_required": {
		"uz": "Kassa idsi to'ldirilmagan",
		"oz": "Касса идси тўлдирилмаган",
		"ru": "Ид кассы не заполнен",
	},

	// RECEIPT
	"send_receipt": {
		"uz": "Chek yuborishda xatolik",
		"oz": "Чек юборишда хатолик",
		"ru": "Ошибка отправки чека",
	},
	"create_receipt": {
		"uz": "Chek yaratishda xatolik",
		"oz": "Чек яратишда хатолик",
		"ru": "Ошибка создания чека",
	},
	"receipt_id_is_required": {
		"uz": "Chek idsi to'ldirilmagan",
		"oz": "Чек идси тўлдирилмаган",
		"ru": "Ид чека не заполнен",
	},
	"user_id_is_required": {
		"uz": "Foydalanuvchi idsi to'ldirilmagan",
		"oz": "Фойдаланувчи идси тўлдирилмаган",
		"ru": "Ид пользователя не заполнен",
	},
	"sale_day_id_is_required": {
		"uz": "Smena idsi to'ldirilmagan",
		"oz": "Смена идси тўлдирилмаган",
		"ru": "Ид смены не заполнен",
	},
	"total_sum_is_required": {
		"uz": "Chek summasi to'ldirilmagan",
		"oz": "Чек суммаси тўлдирилмаган",
		"ru": "Сумма чека не заполнен",
	},
	"total_pay_is_required": {
		"uz": "Chek to'langan summasi to'ldirilmagan",
		"oz": "Чек тўланган суммаси тўлдирилмаган",
		"ru": "Сумма оплаченного чека не заполнен",
	},
	"ofd_fiscal_sign_is_required": {
		"uz": "OFD fiskal imzosi to'ldirilmagan",
		"oz": "OFD фискал имзоси тўлдирилмаган",
		"ru": "Фискальная подпись OFD не заполнен",
	},
	"ofd_terminal_id_is_required": {
		"uz": "OFD terminal idsi to'ldirilmagan",
		"oz": "OFD терминал идси тўлдирилмаган",
		"ru": "Ид терминала OFD не заполнен",
	},
	"ofd_qr_code_url_is_required": {
		"uz": "OFD QR kod urli to'ldirilmagan",
		"oz": "OFD QR код урли тўлдирилмаган",
		"ru": "URL QR кода OFD не заполнен",
	},
	"ofd_date_time_is_required": {
		"uz": "OFD sana va vaqti to'ldirilmagan",
		"oz": "OFD сана ва вақти тўлдирилмаган",
		"ru": "Дата и время OFD не заполнен",
	},
	"ofd_date_time_is_invalid": {
		"uz": "OFD sana va vaqti noto'g'ri",
		"oz": "OFD сана ва вақти нотўғри",
		"ru": "Дата и время OFD неверен",
	},
	"ofd_receipt_sequence_number_is_required": {
		"uz": "OFD chek tartib raqami to'ldirilmagan",
		"oz": "OFD чек тартиб рақами тўлдирилмаган",
		"ru": "Порядковый номер чека OFD не заполнен",
	},
	"created_date_is_required": {
		"uz": "Chek yaratilgan sana va vaqti to'ldirilmagan",
		"oz": "Чек яратилган сана ва вақти тўлдирилмаган",
		"ru": "Дата и время создания чека не заполнен",
	},
	"created_date_is_invalid": {
		"uz": "Chek yaratilgan sana va vaqti noto'g'ri",
		"oz": "Чек яратилган сана ва вақти нотўғри",
		"ru": "Дата и время создания чека неверен",
	},
	"receipt_items_is_required": {
		"uz": "Chek mahsulotlari to'ldirilmagan",
		"oz": "Чек маҳсулотлари тўлдирилмаган",
		"ru": "Продукты чека не заполнен",
	},
	"payment_data_is_required": {
		"uz": "To'lov ma'lumotlari to'ldirilmagan",
		"oz": "Тўлов маълумотлари тўлдирилмаган",
		"ru": "Данные оплаты не заполнен",
	},

	// RECEIPT ITEM
	"create_receipt_item": {
		"uz": "Chek mahsulotini yaratishda xatolik",
		"oz": "Чек маҳсулотини яратишда хатолик",
		"ru": "Ошибка создания продукта чека",
	},
	"receipt_item_id_is_required": {
		"uz": "Chek mahsuloti idsi to'ldirilmagan",
		"oz": "Чек маҳсулоти идси тўлдирилмаган",
		"ru": "Ид продукта чека не заполнен",
	},
	"product_id_is_required": {
		"uz": "Mahsulot idsi to'ldirilmagan",
		"oz": "Маҳсулот идси тўлдирилмаган",
		"ru": "Ид продукта не заполнен",
	},
	"sale_price_is_required": {
		"uz": "Mahsulot sotuv narxi to'ldirilmagan",
		"oz": "Маҳсулот сотув нархи тўлдирилмаган",
		"ru": "Цена продукта не заполнен",
	},
	"sale_count_is_required": {
		"uz": "Mahsulot sotilgan soni to'ldirilmagan",
		"oz": "Маҳсулот сотилган сони тўлдирилмаган",
		"ru": "Количество продукта не заполнен",
	},
	"total_price_is_required": {
		"uz": "Mahsulot umumiy narxi to'ldirilmagan",
		"oz": "Маҳсулот умумий нархи тўлдирилмаган",
		"ru": "Общая цена продукта не заполнен",
	},

	// RECEIPT PAYMENT
	"create_receipt_payment": {
		"uz": "Chek to'lovini yaratishda xatolik",
		"oz": "Чек тўловини яратишда хатолик",
		"ru": "Ошибка создания оплаты чека",
	},
	"payment_data_id_is_required": {
		"uz": "To'lov ma'lumotlari idsi to'ldirilmagan",
		"oz": "Тўлов маълумотлари идси тўлдирилмаган",
		"ru": "Ид данных оплаты не заполнен",
	},
	"payment_type_id_is_required": {
		"uz": "To'lov turi idsi to'ldirilmagan",
		"oz": "Тўлов тури идси тўлдирилмаган",
		"ru": "Ид типа оплаты не заполнен",
	},
	"payment_sum_is_required": {
		"uz": "To'lov summasi to'ldirilmagan",
		"oz": "Тўлов суммаси тўлдирилмаган",
		"ru": "Сумма оплаты не заполнен",
	},

	// SALE DAY
	"create_sale_day": {
		"uz": "Smena yaratishda xatolik",
		"oz": "Смена яратишда хатолик",
		"ru": "Ошибка создания смены",
	},
	"sale_day_already_opened": {
		"uz": "Smena ochilgan",
		"oz": "Смена очилган",
		"ru": "Смена открыта",
	},
	"status_must_be_open": {
		"uz": "Smena ochilmagan",
		"oz": "Смена очилмаган",
		"ru": "Смена не открыта",
	},
	"close_sale_day": {
		"uz": "Smena yopishda xatolik",
		"oz": "Смена ёпишда хатолик",
		"ru": "Ошибка закрытия смены",
	},

	// USER
	"create_user": {
		"uz": "Foydalanuvchi yaratishda xatolik",
		"oz": "Фойдаланувчи яратишда хатолик",
		"ru": "Ошибка создания пользователя",
	},
	"get_user": {
		"uz": "Foydalanuvchi olishda xatolik",
		"oz": "Фойдаланувчи олишда хатолик",
		"ru": "Ошибка получения пользователя",
	},
	"get_user_list": {
		"uz": "Foydalanuvchilar ro'yxatini olishda xatolik",
		"oz": "Фойдаланувчилар рўйхатини олишда хатолик",
		"ru": "Ошибка получения списка пользователей",
	},
	"update_user": {
		"uz": "Foydalanuvchi yangilashda xatolik",
		"oz": "Фойдаланувчи янгилашда хатолик",
		"ru": "Ошибка обновления пользователя",
	},
	"delete_user": {
		"uz": "Foydalanuvchi o'chirishda xatolik",
		"oz": "Фойдаланувчи ўчиришда хатолик",
		"ru": "Ошибка удаления пользователя",
	},
	"reset_password": {
		"uz": "Foydalanuvchi parolini tiklashda xatolik",
		"oz": "Фойдаланувчи паролини тиклашда хатолик",
		"ru": "Ошибка сброса пароля пользователя",
	},

	// USER RELATION
	"add_user_relation": {
		"uz": "Foydalanuvchi relation qo'shishda xatolik",
		"oz": "Фойдаланувчи relation қўшишда хатолик",
		"ru": "Ошибка добавления relation пользователя",
	},
	"remove_user_relation": {
		"uz": "Foydalanuvchi relation o'chirishda xatolik",
		"oz": "Фойдаланувчи relation ўчиришда хатолик",
		"ru": "Ошибка удаления relation пользователя",
	},
	"upsert_user_info": {
		"uz": "Foydalanuvchi ma'lumotlarini yaratishda xatolik",
		"oz": "Фойдаланувчи маълумотларини яратишда хатолик",
		"ru": "Ошибка создания информации пользователя",
	},

	// REPORT
	"sale_day_id_and_company_id_is_required": {
		"uz": "Smena idsi yoki kompaniya idsi to'ldirilmagan",
		"oz": "Смена идси ёки компания идси тўлдирилмаган",
		"ru": "Ид смены или ид компании не заполнен",
	},
	"get_sale_day_info": {
		"uz": "Smena ma'lumotlarini olishda xatolik",
		"oz": "Смена маълумотларини олишда хатолик",
		"ru": "Ошибка получения информации смены",
	},
	"sale_day_info_not_found": {
		"uz": "Smena ma'lumotlari topilmadi",
		"oz": "Смена маълумотлари топилмади",
		"ru": "Информация смены не найдена",
	},
	"company_id_cashbox_id_user_id_is_required": {
		"uz": "Kompaniya idsi yoki kassa idsi yoki foydalanuvchi idsi to'ldirilmagan",
		"oz": "Компания идси ёки касса идси ёки фойдаланувчи идси тўлдирилмаган",
		"ru": "Ид компании или ид кассы или ид пользователя не заполнен",
	},
	"get_receipts_list": {
		"uz": "Cheklar ro'yxatini olishda xatolik",
		"oz": "Чеклар рўйхатини олишда хатолик",
		"ru": "Ошибка получения списка чеков",
	},
	"receipts_list_not_found": {
		"uz": "Cheklar ro'yxati topilmadi",
		"oz": "Чеклар рўйхати топилмади",
		"ru": "Список чеков не найден",
	},
}

var CustomSuccessMessages = map[string]map[string]string{
	// DEFAULT SUCCESS MESSAGE
	"default_success_message": {
		"uz": "Muvaffaqiyatli",
		"oz": "Муваффақиятли",
		"ru": "Успешно",
	},

	// ENTITY
	"create_entity": {
		"uz": "Entity muvaffaqiyatli yaratildi",
		"oz": "Entity муваффақиятли яратилди",
		"ru": "Entity успешно создан",
	},
	"get_entity": {
		"uz": "Entity muvaffaqiyatli olindi",
		"oz": "Entity муваффақиятли олинди",
		"ru": "Entity успешно получен",
	},
	"get_entity_list": {
		"uz": "Entitylar ro'yxati muvaffaqiyatli olindi",
		"oz": "Entityлар рўйхати муваффақиятли олинди",
		"ru": "Список entity успешно получен",
	},
	"update_entity": {
		"uz": "Entity muvaffaqiyatli yangilandi",
		"oz": "Entity муваффақиятли янгиланди",
		"ru": "Entity успешно обновлен",
	},
	"delete_entity": {
		"uz": "Entity muvaffaqiyatli o'chirildi",
		"oz": "Entity муваффақиятли ўчирилди",
		"ru": "Entity успешно удален",
	},
	"join_entity": {
		"uz": "Entity muvaffaqiyatli join qilindi",
		"oz": "Entity муваффақиятли join қилинди",
		"ru": "Entity успешно join",
	},

	// CASHBOX
	"create_cashbox": {
		"uz": "Kassa muvaffaqiyatli yaratildi",
		"oz": "Касса муваффақиятли яратилди",
		"ru": "Касса успешно создана",
	},
	"get_cashbox": {
		"uz": "Kassa muvaffaqiyatli olingan",
		"oz": "Касса муваффақиятли олинган",
		"ru": "Касса успешно получена",
	},
	"get_cashbox_list": {
		"uz": "Kassalar ro'yxati muvaffaqiyatli olingan",
		"oz": "Кассалар рўйхати муваффақиятли олинган",
		"ru": "Список касс успешно получен",
	},
	"update_cashbox": {
		"uz": "Kassa muvaffaqiyatli yangilandi",
		"oz": "Касса муваффақиятли янгиланди",
		"ru": "Касса успешно обновлена",
	},
	"delete_cashbox": {
		"uz": "Kassa muvaffaqiyatli o'chirildi",
		"oz": "Касса муваффақиятли ўчирилди",
		"ru": "Касса успешно удалена",
	},

	// CLIENT PLATFORM
	"create_client_platform": {
		"uz": "Platforma muvaffaqiyatli yaratildi",
		"oz": "Платформа муваффақиятли яратилди",
		"ru": "Платформа успешно создана",
	},
	"get_client_platform": {
		"uz": "Platforma muvaffaqiyatli olingan",
		"oz": "Платформа муваффақиятли олинган",
		"ru": "Платформа успешно получена",
	},
	"get_detailed_client_platform": {
		"uz": "Platforma tafsilotlari muvaffaqiyatli olingan",
		"oz": "Платформа тафсилотлари муваффақиятли олинган",
		"ru": "Подробная информация о платформе успешно получена",
	},
	"get_client_platform_list": {
		"uz": "Platformalar ro'yxati muvaffaqiyatli olingan",
		"oz": "Платформалар рўйхати муваффақиятли олинган",
		"ru": "Список платформ успешно получен",
	},
	"update_client_platform": {
		"uz": "Platforma muvaffaqiyatli yangilandi",
		"oz": "Платформа муваффақиятли янгиланди",
		"ru": "Платформа успешно обновлена",
	},
	"delete_client_platform": {
		"uz": "Platforma muvaffaqiyatli o'chirildi",
		"oz": "Платформа муваффақиятли ўчирилди",
		"ru": "Платформа успешно удалена",
	},

	// CLIENT TYPE
	"create_client_type": {
		"uz": "Klient type muvaffaqiyatli yaratildi",
		"oz": "Клиент типе муваффақиятли яратилди",
		"ru": "Тип клиента успешно создан",
	},
	"get_client_type": {
		"uz": "Klient type muvaffaqiyatli olindi",
		"oz": "Клиент типе муваффақиятли олинди",
		"ru": "Тип клиента успешно получен",
	},
	"get_client_type_list": {
		"uz": "Klient tipi ro'yxati muvaffaqiyatli olindi",
		"oz": "Клиент типи рўйхати муваффақиятли олинди",
		"ru": "Список типов клиентов успешно получен",
	},
	"update_client_type": {
		"uz": "Klient tipi muvaffaqiyatli yangilandi",
		"oz": "Клиент типи муваффақиятли янгиланди",
		"ru": "Тип клиента успешно обновлен",
	},
	"delete_client_type": {
		"uz": "Klient tipi muvaffaqiyatli o'chirildi",
		"oz": "Клиент типи муваффақиятли ўчирилди",
		"ru": "Тип клиента успешно удален",
	},

	// CLIENT
	"add_client": {
		"uz": "Klient muvaffaqiyatli qo'shildi",
		"oz": "Клиент муваффақиятли қўшилди",
		"ru": "Клиент успешно добавлен",
	},
	"get_client_matrix": {
		"uz": "Klient matrixi muvaffaqiyatli olindi",
		"oz": "Клиент матрикси муваффақиятли олинди",
		"ru": "Матрица клиента успешно получена",
	},
	"update_client": {
		"uz": "Klient muvaffaqiyatli yangilandi",
		"oz": "Клиент муваффақиятли янгиланди",
		"ru": "Клиент успешно обновлен",
	},
	"remove_client": {
		"uz": "Klient muvaffaqiyatli o'chirildi",
		"oz": "Клиент муваффақиятли ўчирилди",
		"ru": "Клиент успешно удален",
	},

	// RELATION
	"add_relation": {
		"uz": "Relation muvaffaqiyatli qo'shildi",
		"oz": "Relation муваффақиятли қўшилди",
		"ru": "Relation успешно добавлен",
	},
	"update_relation": {
		"uz": "Relation muvaffaqiyatli yangilandi",
		"oz": "Relation муваффақиятли янгиланди",
		"ru": "Relation успешно обновлен",
	},
	"remove_relation": {
		"uz": "Relation muvaffaqiyatli o'chirildi",
		"oz": "Relation муваффақиятли ўчирилди",
		"ru": "Relation успешно удален",
	},

	// USER INFO
	"add_user_info": {
		"uz": "User info muvaffaqiyatli qo'shildi",
		"oz": "User info муваффақиятли қўшилди",
		"ru": "User info успешно добавлен",
	},
	"update_user_info": {
		"uz": "User info muvaffaqiyatli yangilandi",
		"oz": "User info муваффақиятли янгиланди",
		"ru": "User info успешно обновлен",
	},
	"remove_user_info": {
		"uz": "User info muvaffaqiyatli o'chirildi",
		"oz": "User info муваффақиятли ўчирилди",
		"ru": "User info успешно удален",
	},

	// COMPANY
	"create_company": {
		"uz": "Kompaniya muvaffaqiyatli yaratildi",
		"oz": "Компания муваффақиятли яратилди",
		"ru": "Компания успешно создана",
	},
	"get_company": {
		"uz": "Kompaniya muvaffaqiyatli olindi",
		"oz": "Компания муваффақиятли олинди",
		"ru": "Компания успешно получена",
	},
	"get_company_list": {
		"uz": "Kompaniyalar ro'yxati muvaffaqiyatli olindi",
		"oz": "Компаниялар рўйхати муваффақиятли олинди",
		"ru": "Список компаний успешно получен",
	},
	"update_company": {
		"uz": "Kompaniya muvaffaqiyatli yangilandi",
		"oz": "Компания муваффақиятли янгиланди",
		"ru": "Компания успешно обновлена",
	},
	"delete_company": {
		"uz": "Kompaniya muvaffaqiyatli o'chirildi",
		"oz": "Компания муваффақиятли ўчирилди",
		"ru": "Компания успешно удалена",
	},

	// FIELD TYPE CONFIGURATION
	"get_field_type_configuration": {
		"uz": "Field type configuration muvaffaqiyatli olindi",
		"oz": "Field type configuration муваффақиятли олинди",
		"ru": "Field type configuration успешно получен",
	},

	// GROUP TYPE CONFIGURATION
	"get_group_type_configuration": {
		"uz": "Group type configuration muvaffaqiyatli olindi",
		"oz": "Group type configuration муваффақиятли олинди",
		"ru": "Group type configuration успешно получен",
	},

	// DEFAULT VALUES CONFIGURATION
	"get_default_values_configuration": {
		"uz": "Default values configuration muvaffaqiyatli olindi",
		"oz": "Default values configuration муваффақиятли олинди",
		"ru": "Default values configuration успешно получен",
	},

	// VALIDATION FUNCTION CONFIGURATION
	"get_validation_function_configuration": {
		"uz": "Validation function configuration muvaffaqiyatli olindi",
		"oz": "Validation function configuration муваффақиятли олинди",
		"ru": "Validation function configuration успешно получен",
	},

	// REGEX CONFIGURATION
	"get_regex_configuration": {
		"uz": "Regex configuration muvaffaqiyatli olindi",
		"oz": "Regex configuration муваффақиятли олинди",
		"ru": "Regex configuration успешно получен",
	},

	// FM
	"create_fm": {
		"uz": "FM muvaffaqiyatli yaratildi",
		"oz": "FM муваффақиятли яратилди",
		"ru": "FM успешно создан",
	},
	"get_fm": {
		"uz": "FM muvaffaqiyatli olindi",
		"oz": "FM муваффақиятли олинди",
		"ru": "FM успешно получен",
	},
	"get_fm_list": {
		"uz": "FM ro'yxati muvaffaqiyatli olindi",
		"oz": "FM рўйхати муваффақиятли олинди",
		"ru": "Список FM успешно получен",
	},
	"update_fm": {
		"uz": "FM muvaffaqiyatli yangilandi",
		"oz": "FM муваффақиятли янгиланди",
		"ru": "FM успешно обновлен",
	},
	"delete_fm": {
		"uz": "FM muvaffaqiyatli o'chirildi",
		"oz": "FM муваффақиятли ўчирилди",
		"ru": "FM успешно удален",
	},

	// GROUP
	"create_group": {
		"uz": "Group muvaffaqiyatli yaratildi",
		"oz": "Group муваффақиятли яратилди",
		"ru": "Group успешно создан",
	},
	"get_group": {
		"uz": "Group muvaffaqiyatli olindi",
		"oz": "Group муваффақиятли олинди",
		"ru": "Group успешно получен",
	},
	"get_group_list": {
		"uz": "Group ro'yxati muvaffaqiyatli olindi",
		"oz": "Group рўйхати муваффақиятли олинди",
		"ru": "Список Group успешно получен",
	},
	"update_group": {
		"uz": "Group muvaffaqiyatli yangilandi",
		"oz": "Group муваффақиятли янгиланди",
		"ru": "Group успешно обновлен",
	},
	"delete_group": {
		"uz": "Group muvaffaqiyatli o'chirildi",
		"oz": "Group муваффақиятли ўчирилди",
		"ru": "Group успешно удален",
	},
	"get_full_group": {
		"uz": "Full group muvaffaqiyatli olindi",
		"oz": "Full group муваффақиятли олинди",
		"ru": "Full group успешно получен",
	},

	// ROLE
	"create_role": {
		"uz": "Role muvaffaqiyatli yaratildi",
		"oz": "Role муваффақиятли яратилди",
		"ru": "Role успешно создан",
	},
	"get_role": {
		"uz": "Role muvaffaqiyatli olindi",
		"oz": "Role муваффақиятли олинди",
		"ru": "Role успешно получен",
	},
	"get_role_list": {
		"uz": "Role ro'yxati muvaffaqiyatli olindi",
		"oz": "Role рўйхати муваффақиятли олинди",
		"ru": "Список Role успешно получен",
	},
	"update_role": {
		"uz": "Role muvaffaqiyatli yangilandi",
		"oz": "Role муваффақиятли янгиланди",
		"ru": "Role успешно обновлен",
	},
	"delete_role": {
		"uz": "Role muvaffaqiyatli o'chirildi",
		"oz": "Role муваффақиятли ўчирилди",
		"ru": "Role успешно удален",
	},

	// PERMISSION
	"create_permission": {
		"uz": "Permission muvaffaqiyatli yaratildi",
		"oz": "Permission муваффақиятли яратилди",
		"ru": "Permission успешно создан",
	},
	"get_permission": {
		"uz": "Permission muvaffaqiyatli olindi",
		"oz": "Permission муваффақиятли олинди",
		"ru": "Permission успешно получен",
	},
	"get_permission_list": {
		"uz": "Permission ro'yxati muvaffaqiyatli olindi",
		"oz": "Permission рўйхати муваффақиятли олинди",
		"ru": "Список Permission успешно получен",
	},
	"update_permission": {
		"uz": "Permission muvaffaqiyatli yangilandi",
		"oz": "Permission муваффақиятли янгиланди",
		"ru": "Permission успешно обновлен",
	},
	"delete_permission": {
		"uz": "Permission muvaffaqiyatli o'chirildi",
		"oz": "Permission муваффақиятли ўчирилди",
		"ru": "Permission успешно удален",
	},

	// SCOPE
	"create_user_scope": {
		"uz": "User scope muvaffaqiyatli yaratildi",
		"oz": "User scope муваффақиятли яратилди",
		"ru": "User scope успешно создан",
	},
	"get_scope_list": {
		"uz": "Scope ro'yxati muvaffaqiyatli olindi",
		"oz": "Scope рўйхати муваффақиятли олинди",
		"ru": "Список Scope успешно получен",
	},
	"add_permission_scope": {
		"uz": "Permission scope muvaffaqiyatli qo'shildi",
		"oz": "Permission scope муваффақиятли қўшилди",
		"ru": "Permission scope успешно добавлен",
	},
	"remove_permission_scope": {
		"uz": "Permission scope muvaffaqiyatli o'chirildi",
		"oz": "Permission scope муваффақиятли ўчирилди",
		"ru": "Permission scope успешно удален",
	},

	// ROLE PERMISSION
	"add_role_permission": {
		"uz": "Role permission muvaffaqiyatli qo'shildi",
		"oz": "Role permission муваффақиятли қўшилди",
		"ru": "Role permission успешно добавлен",
	},
	"add_role_permissions": {
		"uz": "Role permissions muvaffaqiyatli qo'shildi",
		"oz": "Role permissions муваффақиятли қўшилди",
		"ru": "Role permissions успешно добавлен",
	},
	"remove_role_permission": {
		"uz": "Role permission muvaffaqiyatli o'chirildi",
		"oz": "Role permission муваффақиятли ўчирилди",
		"ru": "Role permission успешно удален",
	},

	// SESSION
	"login": {
		"uz": "Login muvaffaqiyatli",
		"oz": "Login муваффақиятли",
		"ru": "Успешный логин",
	},
	"logout": {
		"uz": "Logout muvaffaqiyatli",
		"oz": "Logout муваффақиятли",
		"ru": "Успешный логаут",
	},
	"refresh_token": {
		"uz": "Refresh token muvaffaqiyatli",
		"oz": "Refresh token муваффақиятли",
		"ru": "Успешное обновление токена",
	},
	"has_access": {
		"uz": "Has access muvaffaqiyatli",
		"oz": "Has access муваффақиятли",
		"ru": "Успешное получение доступа",
	},

	// PRODUCT
	"add_product": {
		"uz": "Mahsulot muvaffaqiyatli qo'shildi",
		"oz": "Маҳсулот муваффақиятли қўшилди",
		"ru": "Продукт успешно добавлен",
	},
	"update_product": {
		"uz": "Mahsulot muvaffaqiyatli yangilandi",
		"oz": "Маҳсулот муваффақиятли янгиланди",
		"ru": "Продукт успешно обновлен",
	},
	"search_product": {
		"uz": "Mahsulot muvaffaqiyatli qidirildi",
		"oz": "Маҳсулот муваффақиятли қидирилди",
		"ru": "Продукт успешно найден",
	},

	// RECEIPT
	"send_receipt": {
		"uz": "Chek muvaffaqiyatli yuborildi",
		"oz": "Чек муваффақиятли юборилди",
		"ru": "Чек успешно отправлен",
	},

	// SALE DAY
	"create_sale_day": {
		"uz": "Smena muvaffaqiyatli ochildi",
		"oz": "Смена муваффақиятли очилди",
		"ru": "Смена успешно открыта",
	},
	"close_sale_day": {
		"uz": "Smena muvaffaqiyatli yopildi",
		"oz": "Смена муваффақиятли ёпилди",
		"ru": "Смена успешно закрыта",
	},

	// USER
	"create_user": {
		"uz": "Foydalanuvchi muvaffaqiyatli yaratildi",
		"oz": "Фойдаланувчи муваффақиятли яратилди",
		"ru": "Пользователь успешно создан",
	},
	"get_user": {
		"uz": "Foydalanuvchi muvaffaqiyatli olindi",
		"oz": "Фойдаланувчи муваффақиятли олинди",
		"ru": "Пользователь успешно получен",
	},
	"get_user_list": {
		"uz": "Foydalanuvchilar ro'yxati muvaffaqiyatli olindi",
		"oz": "Фойдаланувчилар рўйхати муваффақиятли олинди",
		"ru": "Список пользователей успешно получен",
	},
	"update_user": {
		"uz": "Foydalanuvchi muvaffaqiyatli yangilandi",
		"oz": "Фойдаланувчи муваффақиятли янгиланди",
		"ru": "Пользователь успешно обновлен",
	},
	"delete_user": {
		"uz": "Foydalanuvchi muvaffaqiyatli o'chirildi",
		"oz": "Фойдаланувчи муваффақиятли ўчирилди",
		"ru": "Пользователь успешно удален",
	},
	"reset_password": {
		"uz": "Foydalanuvchi paroli muvaffaqiyatli tiklandi",
		"oz": "Фойдаланувчи пароли муваффақиятли тикланди",
		"ru": "Пароль пользователя успешно сброшен",
	},

	// USER RELATION
	"add_user_relation": {
		"uz": "Foydalanuvchi relation muvaffaqiyatli qo'shildi",
		"oz": "Фойдаланувчи relation муваффақиятли қўшилди",
		"ru": "Пользователь успешно добавлен",
	},
	"remove_user_relation": {
		"uz": "Foydalanuvchi relation muvaffaqiyatli o'chirildi",
		"oz": "Фойдаланувчи relation муваффақиятли ўчирилди",
		"ru": "Пользователь успешно удален",
	},
	"upsert_user_info": {
		"uz": "Foydalanuvchi info muvaffaqiyatli yangilandi",
		"oz": "Фойдаланувчи info муваффақиятли янгиланди",
		"ru": "Пользователь успешно обновлен",
	},

	// REPORT
	"get_sale_day_info": {
		"uz": "Smena ma'lumotlari muvaffaqiyatli olindi",
		"oz": "Смена маълумотлари муваффақиятли олинди",
		"ru": "Информация смены успешно получена",
	},
	"get_receipts_list": {
		"uz": "Cheklar ro'yxati muvaffaqiyatli olindi",
		"oz": "Чеклар рўйхати муваффақиятли олинди",
		"ru": "Список чеков успешно получен",
	},
}
