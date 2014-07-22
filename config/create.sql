drop table trex;
create table if not exists trex (
  id bigint not null auto_increment primary key,
  pipeline  varchar(255),
  pipecount bigint,
  stage     varchar(255),
  stagecount bigint,
  jobname   varchar(255),
  gitrev    varchar(255),
  pass      bool,
  updatetm        timestamp
 );
desc trex;
##insert into trex (pipeline, pipecount, stage, stagecount, jobname, gitrev, pass) values ('pipen',1,'stagen',6,'jobn','12903812098309809821', false);
select * from trex;