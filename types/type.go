package types

type TypeID int

type Type struct {
	id  TypeID
	name string
	Ext  string
	Mime   MIME
}

var (
	nextID TypeID = 0
	types = make(map[TypeID]Type)
	names = make(map[string]TypeID)
	Unknown Type
)

func RegisterType(name string, ext string, mime MIME) Type {
	if _, exists := names[name]; exists { // prevents the creation of multiple Types with the same name (may implement retrieving by name if necessary)
		return types[names[name]]
	}
	id := nextID
	nextID++

	t := Type {
		id: id,
		name: name,
		Ext: ext,
		Mime: mime,
	}

	types[id] = t
	names[name] = id
	return t
}

func init() {
	Unknown = RegisterType("unknown", "", MIME{})
}