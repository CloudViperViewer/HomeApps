
-- Dumping database structure for finance
CREATE DATABASE IF NOT EXISTS `finance` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci */;
USE `finance`;

-- Dumping structure for table finance.fin_accounts
CREATE TABLE IF NOT EXISTS `fin_accounts` (
  `account_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_name` varchar(255) NOT NULL,
  `account_type_id` int(11) NOT NULL,
  `bank_id` int(11) NOT NULL,
  `account_number` varchar(255) NOT NULL,
  `bsb` varchar(6) NOT NULL,
  `balance` double(12,2) NOT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL,
  PRIMARY KEY (`account_id`),
  KEY `fk_bank_id_ref_bank_bank_id` (`bank_id`),
  KEY `index_account_type_id` (`account_type_id`),
  KEY `index_is_active` (`is_active`),
  CONSTRAINT `fk_account_type_id_ref_account_type_account_type_id` FOREIGN KEY (`account_type_id`) REFERENCES `fin_ref_account_type` (`account_type_id`),
  CONSTRAINT `fk_bank_id_ref_bank_bank_id` FOREIGN KEY (`bank_id`) REFERENCES `fin_ref_bank` (`bank_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- Dumping structure for table finance.fin_ref_account_type
CREATE TABLE IF NOT EXISTS `fin_ref_account_type` (
  `account_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`account_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


-- Dumping structure for table finance.fin_ref_bank
CREATE TABLE IF NOT EXISTS `fin_ref_bank` (
  `bank_id` int(11) NOT NULL AUTO_INCREMENT,
  `bank_name` varchar(255) NOT NULL,
  `display_order` int(11) DEFAULT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`bank_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


-- Dumping structure for table finance.fin_ref_transaction_type
CREATE TABLE IF NOT EXISTS `fin_ref_transaction_type` (
  `transaction_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_on` datetime DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`transaction_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
-

-- Dumping structure for table finance.fin_transaction
CREATE TABLE IF NOT EXISTS `fin_transaction` (
  `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL,
  `transaction_type_id` int(11) NOT NULL,
  `value` double(12,2) NOT NULL,
  `recurring_payment_id` int(11) DEFAULT NULL,
  `on_off_bill_id` int(11) DEFAULT NULL,
  `via_paypal` tinyint(1) DEFAULT NULL,
  `date_time` datetime NOT NULL,
  `transaction_with` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL,
  PRIMARY KEY (`transaction_id`),
  KEY `fk_acount_id_ref_accounts_account_id` (`account_id`),
  KEY `fk_trans_type_id` (`transaction_type_id`),
  KEY `index_is_active` (`is_active`),
  CONSTRAINT `fk_acount_id_ref_accounts_account_id` FOREIGN KEY (`account_id`) REFERENCES `fin_accounts` (`account_id`),
  CONSTRAINT `fk_trans_type_id` FOREIGN KEY (`transaction_type_id`) REFERENCES `fin_ref_transaction_type` (`transaction_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;



-- Dumping structure for table finance.fin_vaud_accounts
CREATE TABLE IF NOT EXISTS `fin_vaud_accounts` (
  `audit_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL,
  `action` varchar(255) NOT NULL,
  `column_name` varchar(255) NOT NULL,
  `value_before` varchar(255) DEFAULT NULL,
  `value_after` varchar(255) NOT NULL,
  `user` varchar(255) NOT NULL,
  `date_time` datetime NOT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL,
  PRIMARY KEY (`audit_id`),
  KEY `fk_account_id_fin_accounts_account_id` (`account_id`),
  CONSTRAINT `fk_account_id_fin_accounts_account_id` FOREIGN KEY (`account_id`) REFERENCES `fin_accounts` (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;


-- Dumping structure for table finance.fin_vaud_transactions
CREATE TABLE IF NOT EXISTS `fin_vaud_transactions` (
  `audit_id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_id` int(11) NOT NULL,
  `action` varchar(255) NOT NULL,
  `column_name` varchar(255) NOT NULL,
  `value_before` varchar(255) NOT NULL,
  `value_after` varchar(255) NOT NULL,
  `user` varchar(255) NOT NULL,
  `date_time` datetime NOT NULL,
  `created_by` varchar(255) NOT NULL,
  `created_on` datetime NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_on` datetime DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL,
  PRIMARY KEY (`audit_id`),
  KEY `fk_trans_id_transaction_transaction_id` (`transaction_id`),
  CONSTRAINT `fk_trans_id_transaction_transaction_id` FOREIGN KEY (`transaction_id`) REFERENCES `fin_transaction` (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;



-- Dumping structure for view finance.fin_view_account_type_balance
-- Creating temporary table to overcome VIEW dependency errors
CREATE TABLE `fin_view_account_type_balance` (
	`Account Type` VARCHAR(1) NULL COLLATE 'utf8mb4_uca1400_ai_ci',
	`Account Description` VARCHAR(1) NULL COLLATE 'utf8mb4_uca1400_ai_ci',
	`Total Balance` DOUBLE(19,2) NULL
) ENGINE=MyISAM;

-- Dumping structure for view finance.fin_view_total_balance
-- Creating temporary table to overcome VIEW dependency errors
CREATE TABLE `fin_view_total_balance` (
	`Total Balance` DOUBLE(19,2) NULL
) ENGINE=MyISAM;

-- Dumping structure for trigger finance.fin_accounts_insert_audit
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `fin_accounts_insert_audit` AFTER INSERT ON `fin_accounts` FOR EACH ROW BEGIN 
SET @bank = (SELECT `bank_name` FROM `fin_ref_bank` WHERE `bank_id` = NEW.bank_id);
SET @accountType = (SELECT `name` FROM `fin_ref_account_type` WHERE `account_type_id` = NEW.account_type_id);
SET @isActive = (SELECT IF(NEW.is_active = 1, 'True', 'False'));



END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Dumping structure for trigger finance.fin_accounts_update_audit
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `fin_accounts_update_audit` BEFORE UPDATE ON `fin_accounts` FOR EACH ROW BEGIN 

-- account name
IF OLD.`account_name` != NEW.`account_name` 
THEN
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Account Name', OLD.account_name, NEW.account_name, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- bank
IF OLD.`bank_id` != NEW.`bank_id`
THEN
	SET @newBank = (SELECT `bank_name` FROM `fin_ref_bank` WHERE `bank_id` = NEW.bank_id);
	SET @oldBank = (SELECT `bank_name` FROM `fin_ref_bank` WHERE `bank_id` = OLD.bank_id);
	
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Bank', @oldBank, @newBank, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- account number
IF OLD.`account_number` != NEW.`account_number`
THEN
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Account Number', OLD.account_number, NEW.account_number, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- bsb
IF OLD.`bsb` != NEW.`bsb`
THEN
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'BSB', OLD.bsb, NEW.bsb, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- balance
IF OLD.`balance` != NEW.`balance`
THEN
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Balance', OLD.balance, NEW.balance, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- account type
IF OLD.`account_type_id` != NEW.`account_type_id`	
THEN
	SET @oldAccountType = (SELECT `name` FROM `fin_ref_account_type` WHERE `account_type_id` = OLD.account_type_id);
	SET @newAccountType = (SELECT `name` FROM `fin_ref_account_type` WHERE `account_type_id` = NEW.account_type_id);
	
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Account Type', @oldAccountType, @newAccountType, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;
	
-- is active
IF OLD.`is_active` != NEW.`is_active`	
THEN
	SET @oldIsActive = (SELECT IF(OLD.is_active = 1, 'True', 'False'));
	SET @newIsActive = (SELECT IF(NEW.is_active = 1, 'True', 'False'));
	
	INSERT INTO `fin_vaud_accounts` 
	(account_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(new.account_id, 'UPDATE', 'Is Active', @oldIsActive, @newIsActive, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Dumping structure for trigger finance.fin_transaction_insert_audit
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `fin_transaction_insert_audit`
AFTER INSERT ON `fin_transaction` FOR EACH ROW
BEGIN

SET @taccount = (SELECT CONCAT(acc.account_name, ' ', bank.bank_name) 
						FROM fin_accounts acc 
						LEFT JOIN fin_ref_bank bank ON acc.bank_id = bank.bank_id
						WHERE acc.account_id = NEW.account_id );
SET @ttype = (SELECT `name` FROM `fin_ref_transaction_type` WHERE `transaction_type_id` = NEW.`transaction_type_id`);
SET @isActive = (SELECT IF(NEW.is_Active = 1, 'True', 'False'));


-- recurring payment
IF NEW.recurring_payment_id IS NOT NULL
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'INSERT', 'Recurring Payment', NEW.recurring_payment_id, NEW.created_by, NEW.created_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- on off bull
IF NEW.on_off_bill_id IS NOT NULL
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'INSERT', 'On Off Bill', NEW.on_off_bill_id, NEW.created_by, NEW.created_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- via paypal
IF NEW.via_paypal IS NOT NULL
	THEN
	SET @isPayPal = (SELECT IF(NEW.via_paypal = 1, 'Yes', 'No'));	
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES	
	(NEW.transaction_id,  'INSERT', 'Paypal', @isPayPal, NEW.created_by, NEW.created_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- transaction with
IF NEW.transaction_with IS NOT NULL AND NEW.transaction_with 
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'INSERT', 'Transaction With', NEW.transaction_with, NEW.created_by, NEW.created_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;


END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Dumping structure for trigger finance.fin_transaction_update_audit
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `fin_transaction_update_audit`
BEFORE UPDATE ON `fin_transaction` FOR EACH ROW
BEGIN


-- account
IF NEW.account_id != OLD.account_id
	THEN
	SET @taccountNew = (SELECT CONCAT(acc.account_name, ' ', bank.bank_name) 
						FROM fin_accounts acc 
						LEFT JOIN fin_ref_bank bank ON acc.bank_id = bank.bank_id
						WHERE acc.account_id = NEW.account_id );
	SET @taccountOLD = (SELECT CONCAT(acc.account_name, ' ', bank.bank_name) 
						FROM fin_accounts acc 
						LEFT JOIN fin_ref_bank bank ON acc.bank_id = bank.bank_id
						WHERE acc.account_id = OLD.account_id );
						
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Account',@taccountOLD, @taccountNew, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- transaction type
IF	NEW.`transaction_type_id` != OLD.`transaction_type_id`
	THEN
	SET @ttypeNew = (SELECT `name` FROM `fin_ref_transaction_type` WHERE `transaction_type_id` = NEW.`transaction_type_id`);
	SET @ttypeOld = (SELECT `name` FROM `fin_ref_transaction_type` WHERE `transaction_type_id` = OLD.`transaction_type_id`);
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Transaction Type', @ttypeOld, @ttypeNew, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- value
IF OLD.`value` != NEW.`value`
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Value', OLD.`value`, NEW.`value`, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);	
END IF;

-- date time
IF OLD.date_time != NEW.date_time
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Transaction Date Time', OLD.date_time, NEW.date_time, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- recurring payment
IF NEW.recurring_payment_id IS NOT NULL AND NEW.recurring_payment_id != OLD.recurring_payment_id
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Recurring Payment', OLD.recurring_payment_id, NEW.recurring_payment_id, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- on off bull
IF NEW.on_off_bill_id IS NOT NULL AND NEW.on_off_bill_id != OLD.on_off_bill_id
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'On Off Bill',  OLD.on_off_bill_id, NEW.on_off_bill_id, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- via paypal
IF NEW.via_paypal IS NOT NULL AND NEW.via_paypal != OLD.via_paypal
	THEN
	SET @isPayPalNEW = (SELECT IF(NEW.via_paypal = 1, 'Yes', 'No'));	
	SET @isPayPalOLD= (SELECT IF(OLD.via_paypal = 1, 'Yes', 'No'));
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES	
	(NEW.transaction_id,  'UPDATE', 'Paypal', @isPayPalOLD, @isPayPalNEW, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;

-- transaction with
IF NEW.transaction_with IS NOT NULL AND NEW.transaction_with AND NEW.transaction_with != OLD.transaction_with
	THEN
	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Transaction With', OLD.transaction_with, NEW.transaction_with, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;


-- is active
IF OLD.is_active != NEW.is_active
	THEN
	SET @isActiveNew = (SELECT IF(NEW.is_Active = 1, 'True', 'False'));
	SET @isActiveOLD = (SELECT IF(OLD.is_Active = 1, 'True', 'False'));

	INSERT INTO `fin_vaud_transactions`
	(transaction_id, `action`, `column_name`, value_before, value_after, `user`, date_time, created_by, created_on, updated_by, updated_on, is_active)
	VALUES
	(NEW.transaction_id,  'UPDATE', 'Is Active', @isActiveOLD, @isActiveNew, NEW.updated_by, NEW.updated_on, 'trigger', NOW(), 'trigger', NOW(), 1);
END IF;


END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;

-- Removing temporary table and create final VIEW structure
DROP TABLE IF EXISTS `fin_view_account_type_balance`;
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `fin_view_account_type_balance` AS SELECT accType.`name` AS 'Account Type', accType.`description` AS 'Account Description', SUM(acc.balance) AS 'Total Balance'
FROM `fin_accounts` acc
LEFT JOIN `fin_ref_account_type` accType ON acc.account_type_id = accType.account_type_id AND accType.is_active = 1
WHERE acc.is_active = 1
GROUP BY acc.`account_type_id` 
;

-- Removing temporary table and create final VIEW structure
DROP TABLE IF EXISTS `fin_view_total_balance`;
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `fin_view_total_balance` AS SELECT SUM(`balance`) AS 'Total Balance'
FROM `fin_accounts` WHERE is_active = 1 
;

