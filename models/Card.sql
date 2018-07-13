CREATE TABLE Card (
  CardNum VARCHAR(45) NOT NULL,
  UserId INT NOT NULL,
  Cpassword VARCHAR(45) NOT NULL,
  Brand VARCHAR(45) NOT NULL,
  OverDraft INT NOT NULL DEFAULT 0,
  Remain INT NOT NULL,
  Loan  VARCHAR(45) NOT NULL default "NO",
  LoanBegin VARCHAR(45),
  LoanTime VARCHAR(45),
  Rate VARCHAR(45),
  PRIMARY KEY (CardNum),
  FOREIGN KEY (UserId) REFERENCES User(UserId)on delete cascade on update cascade);

级联操作：
alter table Card add constraint foreign key(UserId) references User(UserId) on delete cascade
on update cascade;
  