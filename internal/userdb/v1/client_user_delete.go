/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package userdb

import userdbentities "github.com/kazekim/UserData/internal/userdb/v1/entities"

func (c *defaultClient) DeleteUser(id int64) error {
	stmt, err := c.db.Prepare("delete from " + userdbentities.TableUser + " where " + userdbentities.TableUserFieldID + " = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
