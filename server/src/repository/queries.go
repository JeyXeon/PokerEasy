package repository

const (
	insertAccountQuery = `INSERT INTO account (account_user_name, money_balance) 
		 				  VALUES ($1, $2)
		 				  ON CONFLICT (account_user_name) DO NOTHING
		 				  RETURNING (account_id, account_user_name, money_balance);`

	selectAccountByIdQuery = `SELECT * FROM account WHERE account_id = $1;`

	updateAccountQuery = `UPDATE account 
						  SET account_user_name = $1, money_balance = $2, connected_lobby_id = $3 
						  WHERE account_id = $4;`

	removeLobbyConnectionFromAccountQuery = `UPDATE account SET connected_lobby_id = null WHERE account_id = $1;`

	insertLobbyQuery = `INSERT INTO lobby(lobby_name, players_amount, creator_id) 
						VALUES ($1, $2, $3) 
						ON CONFLICT (lobby_name) DO NOTHING
						RETURNING (lobby_id, lobby_name, players_amount, creator_id)`

	getLobbyByIdQuery = `SELECT * FROM lobby WHERE lobby_id = $1;`

	getAllLobbiesQuery = `SELECT * FROM lobby`
)
