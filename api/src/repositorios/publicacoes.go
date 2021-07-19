package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorId(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
	select p.*, u.nick from
	publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}
	return publicacao, nil
}
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = ? or s.seguidor_id = ?
	order by p.id desc`,
		usuarioID, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}
	return nil
}
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
	select p.*, u.nick from publicacoes p
	join usuarios u on u.id = p.autor_id
	where p.autor_id = ?`,
		usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}

func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare(`
	update publicacoes set curtidas =
	CASE 
		WHEN curtidas > 0 THEN curtidas - 1
		ELSE 0 
	END
	where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}
