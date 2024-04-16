package lyx_test

import (
	"testing"

	"github.com/pg-sharding/lyx/lyx"

	"github.com/stretchr/testify/assert"
)

func TestSelectComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
				SELECT * FROM "exschema"."extable" WHERE id='83912839012903' AND utype='2' AND btype='sample' AND state = 0 AND is_something = true AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}') AND c_id = 'trunk' ORDER BY entity_id asc;
				`,
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},

				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.AExprOp{
								Left: &lyx.AExprOp{
									Left: &lyx.AExprOp{
										Left: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "id",
											},
											Right: &lyx.AExprSConst{
												Value: "83912839012903",
											},
											Op: "=",
										},
										Right: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "utype",
											},
											Right: &lyx.AExprSConst{
												Value: "2",
											},
											Op: "=",
										},
										Op: "AND",
									},
									Right: &lyx.AExprOp{
										Left: &lyx.ColumnRef{
											ColName: "btype",
										},
										Right: &lyx.AExprSConst{
											Value: "sample",
										},
										Op: "=",
									},
									Op: "AND",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "state",
									},
									Right: &lyx.AExprIConst{
										Value: 0,
									},
									Op: "=",
								},
								Op: "AND",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.ColumnRef{
									ColName: "is_something",
								},
								Right: &lyx.AExprBConst{
									Value: true,
								},
								Op: "=",
							},
							Op: "AND",
						},
						Right: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "keys",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.AExprSConst{
									Value: "{reshke,denchick}",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "keys",
									},
									Right: &lyx.AExprSConst{
										Value: "{munakoiso,werelaxe,x4mmm}",
									},
									Op: "@>",
								},
								Op: "OR",
							},
							Op: "@>",
						},
						Op: "AND",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "c_id",
						},
						Right: &lyx.AExprSConst{
							Value: "trunk",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},

			err: nil,
		},

		{
			query: `
		SELECT *
		FROM "exschema"."extable"
		WHERE
			id='83912839012903'
			AND utype='2'
			AND btype='sample'
			AND state = 0
			AND is_something = true
			AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}')
			AND c_id = 'trunk'
		ORDER BY entity_id asc;
				`,
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.AExprOp{
								Left: &lyx.AExprOp{
									Left: &lyx.AExprOp{
										Left: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "id",
											},
											Right: &lyx.AExprSConst{
												Value: "83912839012903",
											},
											Op: "=",
										},
										Right: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "utype",
											},
											Right: &lyx.AExprSConst{
												Value: "2",
											},
											Op: "=",
										},
										Op: "AND",
									},
									Right: &lyx.AExprOp{
										Left: &lyx.ColumnRef{
											ColName: "btype",
										},
										Right: &lyx.AExprSConst{
											Value: "sample",
										},
										Op: "=",
									},
									Op: "AND",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "state",
									},
									Right: &lyx.AExprIConst{
										Value: 0,
									},
									Op: "=",
								},
								Op: "AND",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.ColumnRef{
									ColName: "is_something",
								},
								Right: &lyx.AExprBConst{
									Value: true,
								},
								Op: "=",
							},
							Op: "AND",
						},
						Right: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "keys",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.AExprSConst{
									Value: "{reshke,denchick}",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "keys",
									},
									Right: &lyx.AExprSConst{
										Value: "{munakoiso,werelaxe,x4mmm}",
									},
									Op: "@>",
								},
								Op: "OR",
							},
							Op: "@>",
						},
						Op: "AND",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "c_id",
						},
						Right: &lyx.AExprSConst{
							Value: "trunk",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
		{
			query: `
				SELECT * 
				FROM a
  				  LEFT JOIN (SELECT * FROM b) c ON a.id = c.id
			`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{&lyx.JoinExpr{
					Larg: &lyx.RangeVar{
						RelationName: "a",
					},
				}},
				TargetList: []lyx.Node{
					&lyx.AExprEmpty{},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: `SELECT a.first, a.second, c.*
					FROM a
  				  	 LEFT JOIN (SELECT * FROM b) c ON a.id = c.id`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{&lyx.JoinExpr{
					Larg: &lyx.RangeVar{
						RelationName: "a",
					},
				}},
				TargetList: []lyx.Node{
					&lyx.ColumnRef{
						ColName:    "first",
						TableAlias: "a",
					},
					&lyx.ColumnRef{
						ColName:    "second",
						TableAlias: "a",
					},
					&lyx.AExprEmpty{},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: `SELECT * FROM a WHERE (SELECT COUNT(*) FROM b) = 1`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "a",
					},
				},
				TargetList: []lyx.Node{
					&lyx.AExprEmpty{},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.Select{
						FromClause: []lyx.FromClauseNode{
							&lyx.RangeVar{RelationName: "b"},
						},
						TargetList: []lyx.Node{
							&lyx.FuncApplication{
								Name: "COUNT",
							},
						},
						Where: &lyx.AExprEmpty{},
					},
					Right: &lyx.AExprIConst{Value: 1},
					Op:    "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectDistinct(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select count(distinct id) from a.b where other_id ~ 1337;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{
					&lyx.FuncApplication{
						Name: "count",
					},
				},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						SchemaName:   "a",
						RelationName: "b",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "other_id",
					},
					Right: &lyx.AExprIConst{
						Value: 1337,
					},
					Op: "~",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSelect(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `SELECT rid FROM shds WHERE sds = '1122' FOR SHARE`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "shds",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "sds",
					},
					Right: &lyx.AExprSConst{
						Value: "1122",
					},
					Op: "=",
				},
				TargetList: []lyx.Node{
					&lyx.ColumnRef{
						ColName: "rid",
					},
				},
			},
			err: nil,
		},
		{
			query: "select 42",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprIConst{Value: 42}},
				Where:      &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 ",
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 limit 7 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i having sum(i)",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i limit 7 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "xx",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprIConst{
						Value: 1,
					},
					Op: "=",
				},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 AND j = 2 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Op: "AND",
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "i",
							TableAlias: "",
						},
						Right: &lyx.AExprIConst{
							Value: 1,
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "",
						},
						Right: &lyx.AExprIConst{
							Value: 2,
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},
		{
			query: `
			select id from tbl where i = 12 for update skip locked limit 1
			`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "tbl",
					},
				},
				Where: &lyx.AExprOp{
					Left:  &lyx.ColumnRef{ColName: "i"},
					Right: &lyx.AExprIConst{Value: 12},
					Op:    "=",
				},
				TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "id"}},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSelectMultipleRelations(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select * from xx, xx2 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
					&lyx.RangeVar{
						RelationName: "xx2",
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "select * from xx, xx2 b where b.i = 1",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
					&lyx.RangeVar{
						RelationName: "xx2",
						Alias:        "b",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "b",
					},
					Right: &lyx.AExprIConst{Value: 1},
					Op:    "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectBetween(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM xxmixed WHERE id BETWEEN 21 AND 22 ORDER BY id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xxmixed",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprList{
						List: []lyx.Node{
							&lyx.AExprIConst{
								Value: 21,
							},
							&lyx.AExprIConst{
								Value: 22,
							},
						},
					},
					Op: "BETWEEN",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSelectAlias(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{

		{
			query: "SELECT kind, sum(len) AS total FROM films GROUP BY kind;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "kind"},
					&lyx.FuncApplication{
						Name: "sum",
						Args: []lyx.Node{
							&lyx.ColumnRef{
								ColName: "len",
							},
						},
					}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "films",
				},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT kind, sum(len) AS total FROM films AS f GROUP BY kind;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "kind"},
					&lyx.FuncApplication{
						Name: "sum",
						Args: []lyx.Node{
							&lyx.ColumnRef{
								ColName: "len",
							},
						},
					}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "films",
					Alias:        "f",
				},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "create table xx ( i int )",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		/* same test, no spaces between tokens */
		{
			query: "create table xx(i int)",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		/* same test, no spaces between tokens */
		{
			query: "create table tt(i int);",
			exp: &lyx.CreateTable{
				TableName: "tt",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE sshjt1(i int, j int);",
			exp: &lyx.CreateTable{
				TableName: "sshjt1",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
					{
						ColName: "j",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE orders(id INT PRIMARY KEY);",
			exp: &lyx.CreateTable{
				TableName: "orders",
				TableElts: []lyx.TableElt{
					{
						ColName: "id",
						ColType: "INT",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE delivery(id INT PRIMARY KEY, order_id INT, FOREIGN KEY(order_id) REFERENCES orders(id));",
			exp: &lyx.CreateTable{
				TableName: "delivery",
				TableElts: []lyx.TableElt{
					{
						ColName: "id",
						ColType: "INT",
					},
					{
						ColName: "order_id",
						ColType: "INT",
					},
				},
			},
			err: nil,
		},

		{
			query: "create database reg",
			exp:   &lyx.CreateDatabase{},
			err:   nil,
		},

		{
			query: "create role reg",
			exp:   &lyx.CreateRole{},
			err:   nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "insert into xx (id) values(1)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"id"},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{&lyx.AExprIConst{Value: 1}},
				},
			},
			err: nil,
		},

		{
			query: "INSERT INTO xx (w_id) SELECT 20;",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"w_id"},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprIConst{Value: 20}},
					Where:      &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},

		{
			query: "Insert into xx (i) select * from yy where i = 8",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"i"},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{
						&lyx.RangeVar{
							RelationName: "yy",
						},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "i",
						},
						Right: &lyx.AExprIConst{Value: 8},
						Op:    "=",
					},
				},
			},
			err: nil,
		},

		{
			query: `INSERT INTO films (code, title, did, date_prod, kind)
			VALUES ('T_601', 'Yojimbo', 106, '1961-06-16', 'Drama');`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"code", "title", "did", "date_prod", "kind"},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprSConst{Value: "T_601"},
						&lyx.AExprSConst{Value: "Yojimbo"},
						&lyx.AExprIConst{Value: 106},
						&lyx.AExprSConst{Value: "1961-06-16"},
						&lyx.AExprSConst{Value: "Drama"},
					},
				},
			},
			err: nil,
		},

		///

		{
			query: `INSERT INTO films SELECT * FROM tmp_films WHERE date_prod < '2004-05-07';`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "tmp_films",
					},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "date_prod",
							TableAlias: "",
						},
						Right: &lyx.AExprSConst{
							Value: "2004-05-07",
						},
						Op: "<",
					},
				},
			},
			err: nil,
		},

		{
			query: "insert into xx (id,id2) values(1,2)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{
					"id",
					"id2",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{Value: 1},
						&lyx.AExprIConst{Value: 2},
					},
				},
			},
			err: nil,
		},
		{
			query: "INSERT INTO xxtt1 (j, i, w_id) VALUES(2121221, -211212, 21);",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xxtt1",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"j", "i", "w_id"},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{Value: 2121221},
						&lyx.AExprIConst{Value: -211212},
						&lyx.AExprIConst{Value: 21},
					},
				},
			},
			err: nil,
		},

		/* only first tuple from values clause parsed  */
		{
			query: "insert into xx (id,id2) values(1,2), (2,3), ( 4, 5)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{
					"id",
					"id2",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{Value: 1},
						&lyx.AExprIConst{Value: 2},
					},
				},
			},
			err: nil,
		},

		/* insert from select */
		{
			query: "insert into xx select * from xx2 where id2 = 7",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "xx2",
					},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "id2",
						},
						Right: &lyx.AExprIConst{
							Value: 7,
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},

		/* Explicit row */
		{
			query: `
			INSERT INTO lol.kek
(
 id,
 info
 )
VALUES(2822, ROW('none', '0'));
`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "kek",
					SchemaName:   "lol",
					Alias:        "",
				},

				Columns: []string{"id", "info"},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{
							Value: 2822,
						},
						nil,
					},
				},
			},
			err: nil,
		},
		{
			query: `
			INSERT INTO lol.kek
(
 id,
 info,
 bytes
 )
VALUES(2822, ROW('none', '0'), ''::bytea);
`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "kek",
					SchemaName:   "lol",
					Alias:        "",
				},
				Columns: []string{"id", "info", "bytes"},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{
							Value: 2822,
						},
						nil,
						&lyx.AExprSConst{
							Value: "",
						},
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestInsertComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			INSERT INTO
			cluster_membership(membership_partition, host_id, rpc_address, rpc_port, role, session_start, last_heartbeat, record_expiry)
			VALUES('0', '\x0d1e9794bac311ee8a19ea2e156ae3d8', '2a02:6b8:c08:6819:0:5c2f:8548:0', '6934', '2', '2024-01-24 14:15:39.699192Z', '2024-01-24 14:15:39.750655Z', '2024-01-26 14:15:39.750655Z')
			ON CONFLICT(membership_partition, host_id)
			DO UPDATE SET 
			membership_partition = '0', host_id = '\x0d1e9794bac311ee8a19ea2e156ae3d8',
			rpc_address = '2a02:6b8:c08:6819:0:5c2f:8548:0', rpc_port = '6934', role = '2', 
			session_start = '2024-01-24 14:15:39.699192Z', last_heartbeat = '2024-01-24 14:15:39.750655Z', record_expiry = '2024-01-26 14:15:39.750655Z'
			`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "cluster_membership",
				},
				Columns: []string{
					"membership_partition",
					"host_id",
					"rpc_address",
					"rpc_port",
					"role",
					"session_start",
					"last_heartbeat",
					"record_expiry",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprSConst{
							Value: "0",
						},
						&lyx.AExprSConst{
							Value: "\\x0d1e9794bac311ee8a19ea2e156ae3d8",
						},
						&lyx.AExprSConst{
							Value: "2a02:6b8:c08:6819:0:5c2f:8548:0",
						},
						&lyx.AExprSConst{
							Value: "6934",
						},
						&lyx.AExprSConst{
							Value: "2",
						},
						&lyx.AExprSConst{
							Value: "2024-01-24 14:15:39.699192Z",
						},
						&lyx.AExprSConst{
							Value: "2024-01-24 14:15:39.750655Z",
						},
						&lyx.AExprSConst{
							Value: "2024-01-26 14:15:39.750655Z",
						},
					},
				},
			},
			err: nil,
		},
		{
			query: `INSERT INTO "people" ("first_name","last_name","email","id") VALUES ('John','Smith','','1') RETURNING "i"`,
			err:   nil,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "people",
				},
				Columns: []string{
					"first_name",
					"last_name",
					"email",
					"id",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprSConst{
							Value: "John",
						},
						&lyx.AExprSConst{
							Value: "Smith",
						},
						&lyx.AExprSConst{Value: ""},
						&lyx.AExprSConst{Value: "1"},
					},
				},
			},
		},
		{
			query: `

INSERT
	INTO "exschema"."extable"
("cat","fox","dog","fly","bee","lol","kek","type","id","version","isok","isnotok","keys","tags","key_id",
	"key_value","value_type","state")
VALUES ('1970-01-01 12:00:00.5',111111,NULL,NULL,'9223372036854775807',
	'2','some','thing','*()*()Q*D()beiwe','0','trunk',true,'{280,fb8,909,e6,ffc}',
	'{9223372036854775806}','31337','bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././','0',0)

			`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
					Alias:        "",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprSConst{Value: "1970-01-01 12:00:00.5"},
						&lyx.AExprIConst{Value: 111111},
						&lyx.AExprNConst{},
						&lyx.AExprNConst{},
						&lyx.AExprSConst{Value: "9223372036854775807"},
						&lyx.AExprSConst{Value: "2"},
						&lyx.AExprSConst{Value: "some"},
						&lyx.AExprSConst{Value: "thing"},
						&lyx.AExprSConst{Value: "*()*()Q*D()beiwe"},
						&lyx.AExprSConst{Value: "0"},
						&lyx.AExprSConst{Value: "trunk"},
						&lyx.AExprBConst{Value: true},
						&lyx.AExprSConst{Value: "{280,fb8,909,e6,ffc}"},
						&lyx.AExprSConst{Value: "{9223372036854775806}"},
						&lyx.AExprSConst{Value: "31337"},
						&lyx.AExprSConst{Value: "bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././"},
						&lyx.AExprSConst{Value: "0"},
						&lyx.AExprIConst{Value: 0},
					},
				},
				Columns: []string{
					"cat",
					"fox",
					"dog",
					"fly",
					"bee",
					"lol",
					"kek",
					"type",
					"id",
					"version",
					"isok",
					"isnotok",
					"keys",
					"tags",
					"key_id",
					"key_value",
					"value_type",
					"state",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* simple update */
		{
			query: "UPDATE films SET kind = 'Dramatic' WHERE kind = 'Drama';",
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Right: &lyx.AExprSConst{Value: "Drama"},
					Op:    "=",
				},
			},
			err: nil,
		},

		/* another simple update */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03';`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Right: &lyx.AExprSConst{
							Value: "San Francisco",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Right: &lyx.AExprSConst{
							Value: "2003-07-03",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},

		/* another simple update  with returning clause */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03'
			RETURNING temp_lo, temp_hi, prcp;`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Right: &lyx.AExprSConst{
							Value: "San Francisco",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Right: &lyx.AExprSConst{
							Value: "2003-07-03",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* delete without where */
		{
			query: "DELETE FROM films;",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		/* simple delete */
		{
			query: "DELETE FROM films WHERE kind <> 'Musical';",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Right: &lyx.AExprSConst{
						Value: "Musical",
					},
					Op: "<>",
				},
			},
			err: nil,
		},

		/* simple delete with returning */
		{
			query: "DELETE FROM tasks WHERE status = 'DONE' RETURNING *;",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "tasks",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "status",
					},
					Right: &lyx.AExprSConst{
						Value: "DONE",
					},
					Op: "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* copy with where */
		{
			query: "COPY (SELECT * FROM country WHERE country_name = 'R') TO  STDOUT",
			exp: &lyx.Copy{
				IsFrom: false,
				SubStmt: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.AExprEmpty{},
					},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "country",
					}},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "country_name",
						},
						Right: &lyx.AExprSConst{
							Value: "R",
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},
		/* copy with where */
		{
			query: "COPY copy_test FROM STDIN WHERE id <= 30;",
			exp: &lyx.Copy{
				TableRef: &lyx.RangeVar{
					RelationName: "copy_test",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprIConst{
						Value: 30,
					},
					Op: "<=",
				},
				IsFrom: true,
			},
			err: nil,
		},
		{
			query: `copy pgbench_accounts from stdin with (freeze on)`,
			exp: &lyx.Copy{
				TableRef: &lyx.RangeVar{
					RelationName: "pgbench_accounts",
				},
				Where:  &lyx.AExprEmpty{},
				IsFrom: true,
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestFuncApplication(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}
	for _, tt := range []tcase{
		{
			query: "INSERT INTO xxtt1 (j, w_id) SELECT a, 20 from unnest(ARRAY[10]) a;",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{RelationName: "xxtt1"},
				Columns:  []string{"j", "w_id"},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.ColumnRef{
							ColName: "a",
						},
						&lyx.AExprIConst{
							Value: 20,
						},
					},
					FromClause: []lyx.FromClauseNode{nil},
					Where:      &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},

		{
			query: ` UPDATE xxtt1 
				set i=a.i, j=a.j 
			from unnest(ARRAY[(1,10)]) as a(i int, j int) 
			where w_id=20 and xxtt1.j=a.j;`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{RelationName: "xxtt1"},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "w_id",
						},
						Right: &lyx.AExprIConst{
							Value: 20,
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "xxtt1",
							ColName:    "j",
						},
						Right: &lyx.ColumnRef{
							TableAlias: "a",
							ColName:    "j",
						},
						Op: "=",
					},
					Op: "and",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSelectTargetLists(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{

		{
			query: `
			SELECT NOT EXISTS (SELECT 1) AS invalid
			`,

			exp: &lyx.Select{
				TargetList: []lyx.Node{
					nil,
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: `SELECT current_schema()`,
			exp: &lyx.Select{
				Where: &lyx.AExprEmpty{},
				TargetList: []lyx.Node{
					&lyx.FuncApplication{
						Name: "current_schema",
					},
				},
			},
			err: nil,
		},
		{
			query: "SELECT pg_is_in_recovery(), id FROM tsa_test WHERE id = 22;",
			exp: &lyx.Select{

				TargetList: []lyx.Node{
					&lyx.FuncApplication{
						Name: "pg_is_in_recovery",
					},
					&lyx.ColumnRef{ColName: "id"},
				},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "tsa_test",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprIConst{
						Value: 22,
					},
					Op: "=",
				},
			},
			err: nil,
		},
		{
			query: "select count(*) from pgbench_branches;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{
					&lyx.FuncApplication{
						Name: "count",
					},
				},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "pgbench_branches",
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestWithComments(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * /* o lo lo l ol **** *!*&*(&!89**/ FROM  tsa_test WHERE id = 22;",
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "tsa_test",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprIConst{
						Value: 22,
					},
					Op: "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestJoins(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM delivery LEFT JOIN orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT * FROM sshjt1 a join sshjt1 b ON TRUE WHERE a.i = 12 AND b.j = a.j;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "sshjt1",
							Alias:        "a",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "sshjt1",
							Alias:        "b",
						},
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "i",
							TableAlias: "a",
						},
						Right: &lyx.AExprIConst{
							Value: 12,
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "b",
						},
						Right: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "a",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestMisc(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "vacuum xx;",
			exp:   &lyx.Vacuum{},
			err:   nil,
		},
		{
			query: "drop table xx;",
			exp:   &lyx.Drop{},
			err:   nil,
		},
		{
			query: "analyze xx;",
			exp:   &lyx.Analyze{},
			err:   nil,
		},
		{
			query: "alter table xx add column i int;",
			exp:   &lyx.Alter{},
			err:   nil,
		},
		{
			query: "cluster xx;",
			exp:   &lyx.Cluster{},
			err:   nil,
		},
		{
			query: ";",
			exp:   nil,
			err:   nil,
		},
		{
			query: "-- comment",
			exp:   nil,
			err:   nil,
		},
		{
			query: "/* comment */",
			exp:   nil,
			err:   nil,
		},

		{
			query: "select 11 /* comment */",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprIConst{Value: 11}},
				Where:      &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "select 0.017264",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprSConst{Value: "0.017264"}},
				Where:      &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT  * FROM delivery  JOIN  /* comment */ orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT  * FROM delivery  JOIN  /* comment */ orders ON order_id = id;-- ololol",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT /* comment */  * FROM delivery  JOIN  orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "begin;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_BEGIN,
			},
			err: nil,
		},
		{
			query: "begin transaction read only;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_BEGIN,
				Options: []lyx.TransactionModeItem{
					lyx.TransactionReadOnly,
				},
			},
			err: nil,
		},
		{
			query: "rollback;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_ROLLBACK,
			},
			err: nil,
		},
		{
			query: "commit;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_COMMIT,
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestErrors(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM sshjt1 a join sshjt1 b WHERE a.i = 12 ON TRUE;",
			exp:   nil,
			err:   nil,
		},
		{
			query: "SELECT * FROM x WHERE ixxxd = 1 iuwehiuhweui;",
			exp:   nil,
			err:   nil,
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.Error(err, "query %s", tt.query)
	}
}

func TestOperators(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT 1 ~~~~~~~~ 'dewoijoi'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprOp{
					Left:  &lyx.AExprIConst{Value: 1},
					Right: &lyx.AExprSConst{Value: "dewoijoi"},
					Op:    "~~~~~~~~",
				}},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT 1 ~~~^%^%^^~~~~~ 'dewoijoi'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprOp{
					Left:  &lyx.AExprIConst{Value: 1},
					Right: &lyx.AExprSConst{Value: "dewoijoi"},
					Op:    "~~~^%^%^^~~~~~",
				}},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM xxtt1 a WHERE a.w_id = 21 and j + i != 0;'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "xxtt1",
						Alias:        "a",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "w_id",
							TableAlias: "a",
						},
						Right: &lyx.AExprIConst{
							Value: 21,
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "j",
							},
							Right: &lyx.ColumnRef{
								ColName: "i",
							},
							Op: "+",
						},
						Right: &lyx.AExprIConst{
							Value: 0,
						},
						Op: "!=",
					},
					Op: "and",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestParams(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "PREPARE P(TEXT) AS SELECT $1",
			exp: &lyx.PrepareStmt{
				Name: "P",
				Statement: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.ParamRef{
							Number: 1,
						},
					},
					Where: &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM x WHERE i = $1",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{RelationName: "x"},
				},
				Where: &lyx.AExprOp{
					Op:    "=",
					Left:  &lyx.ColumnRef{ColName: "i"},
					Right: &lyx.ParamRef{Number: 1},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreateTableWithPrimaryKey(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			create table IF NOT EXISTS item (i_id int not null, i_im_id int, i_name varchar(24), i_price decimal(5,2), i_data varchar(50),PRIMARY KEY(i_id) )`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestCreateTableWithCompositePrimaryKey(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			create table IF NOT EXISTS customer (c_id int not null, c_d_id int not null, c_w_id int not null, c_first varchar(16), c_middle char(2), c_last varchar(16), c_street_1 varchar(20), c_street_2 varchar(20), c_city varchar(20), c_state char(2), c_zip char(9), c_phone char(16), c_credit char(2), c_credit_lim bigint, c_discount decimal(4,2), c_balance decimal(12,2), c_ytd_payment decimal(12,2), c_payment_cnt int, c_delivery_cnt int, c_data text, PRIMARY KEY(c_w_id, c_d_id, c_id) )`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestDrop(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			drop table if exists pgbench_accounts, pgbench_branches, pgbench_history, pgbench_tellers`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestCreateFail(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
	}

	for _, tt := range []tcase{
		{
			query: "create table SELECT ( i int )",
		},
		{
			query: "create table ASYMMETRIC ( i int )",
		},
		{
			query: "create table INTEGER ( i int )",
		},
		{
			query: "create table xx ( i SELECT )",
		},
		{
			query: "create table xx ( i WHERE )",
		},
		{
			query: "create table xx ( i AS )",
		}} {
		_, err := lyx.Parse(tt.query)
		assert.Error(err, tt.query)
	}
}

func TestTruncateTable(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}
	for _, tt := range []tcase{
		{
			query: `
			truncate table pgbench_accounts, pgbench_branches, pgbench_history, pgbench_tellers
			`,

			exp: &lyx.Truncate{},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestCreateTableWith(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			create table 
				pgbench_tellers(tid int not null,bid int,tbalance int,filler char(84)) with (fillfactor=100)
			`,
			exp: &lyx.CreateTable{
				TableName: "pgbench_tellers",
				TableElts: []lyx.TableElt{
					{
						ColName: "tid",
						ColType: "int",
					},
					{
						ColName: "bid",
						ColType: "int",
					},
					{
						ColName: "tbalance",
						ColType: "int",
					},
					{
						ColName: "filler",
						ColType: "char",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestCreateSuccess(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "create table xx ( ADMIN int, ATOMIC int, CLASS int, LIKE int )",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "ADMIN",
						ColType: "int",
					},
					{
						ColName: "ATOMIC",
						ColType: "int",
					},
					{
						ColName: "CLASS",
						ColType: "int",
					},
					{
						ColName: "LIKE",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table xx ( i0 BINARY, i1 CURRENT_SCHEMA, i2 IS, i3 JOIN, i4 NOTNULL, i5 TABLESAMPLE)",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i0",
						ColType: "BINARY",
					},
					{
						ColName: "i1",
						ColType: "CURRENT_SCHEMA",
					},
					{
						ColName: "i2",
						ColType: "IS",
					},
					{
						ColName: "i3",
						ColType: "JOIN",
					},
					{
						ColName: "i4",
						ColType: "NOTNULL",
					},
					{
						ColName: "i5",
						ColType: "TABLESAMPLE",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table JSON_ARRAYAGG ( i int )",
			exp: &lyx.CreateTable{
				TableName: "JSON_ARRAYAGG",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table XMLTABLE ( i int )",
			exp: &lyx.CreateTable{
				TableName: "XMLTABLE",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		}} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestCte(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{

		{
			query: `
with cte (i) as (values (12), (13))
select * from tbl inner join cte on tbl.i = cte.i;
			`,
			exp: &lyx.Select{
				WithClause: []lyx.Node{
					&lyx.ValueClause{
						Values: []lyx.Node{
							&lyx.AExprIConst{
								Value: 12,
							},
						},
					},
				},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "tbl",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "cte",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
				TargetList: []lyx.Node{
					&lyx.AExprEmpty{},
				},
			},
			err: nil,
		},

		{
			query: `
			WITH xxxx AS (
				SELECT * from t where i = 1
			) 
			SELECT * from xxxx WHERE i = 2;
			`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "xxxx",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "i",
					},
					Right: &lyx.AExprIConst{Value: 2},
					Op:    "=",
				},
				TargetList: []lyx.Node{
					&lyx.AExprEmpty{},
				},

				WithClause: []lyx.Node{
					&lyx.Select{
						FromClause: []lyx.FromClauseNode{
							&lyx.RangeVar{
								RelationName: "t",
							},
						},
						Where: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "i",
							},
							Right: &lyx.AExprIConst{Value: 1},
							Op:    "=",
						},
						TargetList: []lyx.Node{
							&lyx.AExprEmpty{},
						},
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestMiscQ(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SELECT * from X where id in ( select 1 );
			`,
			exp: &lyx.Select{
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "X",
					},
				},
				TargetList: []lyx.Node{
					&lyx.AExprEmpty{},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},

					Right: &lyx.Select{
						Where: &lyx.AExprEmpty{},
						TargetList: []lyx.Node{
							&lyx.AExprIConst{Value: 1},
						},
					},
					Op: "IN",
				},
			},
			err: nil,
		},

		{
			query: `
			delete from tbl where id in (select id from tbl where i = 12 for update skip locked limit 1)
			and i = 12 returning *;
			`,
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "tbl",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{

						Left: &lyx.ColumnRef{
							ColName: "id",
						},
						Right: &lyx.Select{
							FromClause: []lyx.FromClauseNode{
								&lyx.RangeVar{
									RelationName: "tbl",
								},
							},
							Where: &lyx.AExprOp{
								Left:  &lyx.ColumnRef{ColName: "i"},
								Right: &lyx.AExprIConst{Value: 12},
								Op:    "=",
							},
							TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "id"}},
						},
						Op: "IN",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "i",
						},
						Right: &lyx.AExprIConst{
							Value: 12,
						},
						Op: "=",
					},
					Op: "and",
				},
			},
			err: nil,
		},
		{
			query: `
		 DELETE FROM
		 	cl_m 
			 WHERE s_id = 
			 ANY(ARRAY(SELECT s_id 
				FROM cl_m 
				WHERE mem_par = '0' 
				AND r_e < '2024-01-23 10:18:52.937505Z'))
		 `,
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "cl_m",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "s_id",
					},
					Right: &lyx.FuncApplication{
						Name: "ANY",
						Args: []lyx.Node{
							&lyx.SubLink{
								SubSelect: &lyx.Select{
									FromClause: []lyx.FromClauseNode{
										&lyx.RangeVar{
											RelationName: "cl_m",
										},
									},
									TargetList: []lyx.Node{
										&lyx.ColumnRef{
											ColName: "s_id",
										},
									},
									Where: &lyx.AExprOp{
										Left: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "mem_par",
											},
											Right: &lyx.AExprSConst{
												Value: "0",
											},
											Op: "=",
										},
										Right: &lyx.AExprOp{

											Left: &lyx.ColumnRef{
												ColName: "r_e",
											},
											Right: &lyx.AExprSConst{
												Value: "2024-01-23 10:18:52.937505Z",
											},
											Op: "<",
										},
										Op: "AND",
									},
								},
							},
						},
					},
					Op: "=",
				},
			},
			err: nil,
		},
		// {
		// 	query: `
		// SELECT n.nspname as "Schema",
		//   c.relname as "Name",
		//   CASE c.relkind WHEN 'r' THEN 'table' WHEN 'v' THEN 'view' WHEN 'm' THEN 'materialized view' WHEN 'i' THEN 'index' WHEN 'S' THEN 'sequence' WHEN 't' THEN 'TOAST table' WHEN 'f' THEN 'foreign table' WHEN 'p' THEN 'partitioned table' WHEN 'I' THEN 'partitioned index' END as "Type",
		//   pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
		// FROM pg_catalog.pg_class c
		//      LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
		//      LEFT JOIN pg_catalog.pg_am am ON am.oid = c.relam
		// WHERE c.relkind IN ('r','p','v','m','S','f','')
		//       AND n.nspname <> 'pg_catalog'
		//       AND n.nspname !~ '^pg_toast'
		//       AND n.nspname <> 'information_schema'
		//   AND pg_catalog.pg_table_is_visible(c.oid)
		// ORDER BY 1,2;
		// `,
		// 	err: nil,
		// },
		{
			query: `
		INSERT INTO warehouse1
				(w_id, w_name, w_street_1, w_street_2, w_city, w_state, w_zip, w_tax, w_ytd)
		VALUES (1, 'name-vjxqu','street1-qkfzdggwut','street2-jxuhvhtqct', 'city-irchbmwruo', 'er', 'zip-26599', 0.017264,300000);'
					`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					SchemaName:   "",
					RelationName: "warehouse1",
				},
				Columns: []string{
					"w_id",
					"w_name",
					"w_street_1",
					"w_street_2",
					"w_city",
					"w_state",
					"w_zip",
					"w_tax",
					"w_ytd",
				},
				SubSelect: &lyx.ValueClause{
					Values: []lyx.Node{
						&lyx.AExprIConst{
							Value: 1,
						},
						&lyx.AExprSConst{
							Value: "name-vjxqu",
						},
						&lyx.AExprSConst{
							Value: "street1-qkfzdggwut",
						},
						&lyx.AExprSConst{
							Value: "street2-jxuhvhtqct",
						},
						&lyx.AExprSConst{
							Value: "city-irchbmwruo",
						},
						&lyx.AExprSConst{
							Value: "er",
						},
						&lyx.AExprSConst{
							"zip-26599",
						},
						&lyx.AExprSConst{
							"0.017264",
						},
						&lyx.AExprIConst{
							300000,
						},
					},
				},
			},
			err: nil,
		},
		{
			query: `
					SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL READ COMMITTED
					`,
			exp: &lyx.VariableSetStmt{
				Session: false,
				IsLocal: false,
				Default: false,
				Name:    "",
				Value:   nil,
				Kind:    "SET",
				TxMode: []lyx.TransactionModeItem{
					0,
				},
			},
			err: nil,
		},
		{
			query: `
					select ttt.i from ttt;
					`,
			exp: &lyx.Select{
				TargetList: []lyx.Node{
					&lyx.ColumnRef{
						TableAlias: "ttt",
						ColName:    "i",
					},
				},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "ttt",
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: `
					UPDATE customer1
					    SET c_balance= -4755.000000, c_ytd_payment=4755.000000
					WHERE c_w_id = 1
					AND c_d_id=1
					AND c_id=963
					`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "customer1",
				},
				Where: &lyx.AExprOp{
					Op: "AND",
					Left: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "c_w_id",
							},
							Right: &lyx.AExprIConst{
								Value: 1,
							},
							Op: "=",
						},
						Right: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "c_d_id",
							},
							Right: &lyx.AExprIConst{
								Value: 1,
							},
							Op: "=",
						},
						Op: "AND",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "c_id",
						},
						Right: &lyx.AExprIConst{
							Value: 963,
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestExplain(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			EXPLAIN select 1
`,
			exp: &lyx.Explain{
				Stmt: &lyx.Select{
					Where: &lyx.AExprEmpty{},
					TargetList: []lyx.Node{
						&lyx.AExprIConst{
							Value: 1,
						},
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSetStmt(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SET search_path to 'test'
`,
			exp: &lyx.VariableSetStmt{
				Session: false,
				Name:    "search_path",
				Kind:    lyx.VarTypeSet,
				Value: []string{
					"test",
				},
			},
			err: nil,
		},
		{
			query: `
			SET extra_float_digits = 3
`,
			exp: &lyx.VariableSetStmt{
				Session: false,
				Name:    "extra_float_digits",
				Kind:    lyx.VarTypeSet,
				Value: []string{
					"3",
				},
			},
			err: nil,
		},
		{
			query: `
			SET search_path = 'test'
`,
			exp: &lyx.VariableSetStmt{
				Session: false,
				Name:    "search_path",
				Kind:    lyx.VarTypeSet,
				Value: []string{
					"test",
				},
			},
			err: nil,
		},
		{
			query: `
			set __spqr_lol to 'kek';`,
			exp: &lyx.VariableSetStmt{
				Session: false,
				Name:    "__spqr_lol",
				Kind:    lyx.VarTypeSet,
				Value: []string{
					"kek",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestShowStmt(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SHOW transaction_read_only
`,
			exp: &lyx.VariableShowStmt{
				Name: "transaction_read_only",
			},
			err: nil,
		},
		{
			query: `
			SHOW all
`,
			exp: &lyx.VariableShowStmt{
				Name: "all",
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSetSessionStmt(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SET SESSION CHARACTERISTICS AS TRANSACTION READ ONLY`,
			exp: &lyx.VariableSetStmt{
				Kind: lyx.VarTypeSet,
				TxMode: []lyx.TransactionModeItem{
					lyx.TransactionReadOnly,
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}
