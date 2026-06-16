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
	Types = make(map[TypeID]Type)
	Names = make(map[string]TypeID)
	Unknown Type
)

func RegisterType(name string, ext string, mime MIME) Type {
	// prevents the creation of multiple Types with the same name (may implement retrieving by name if necessary)
	if _, exists := Names[name]; exists {
		return Types[Names[name]]
	}
	id := nextID
	nextID++

	t := Type {
		id: id,
		name: name,
		Ext: ext,
		Mime: mime,
	}

	Types[id] = t
	Names[name] = id
	return t
}

func init() {
	Unknown = RegisterType("unknown", "", MIME{})
}