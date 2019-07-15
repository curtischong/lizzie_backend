create schema typer;
create schema lnews;
create schema bio;
create schema emotion;
create schema life_event;
create schema peaks;

-- Typer

create table typer.text_box (
  unixt bigint not null, -- default (extract(epoch from now())*1000)::bigint,
  ts timestamp not null, -- without time zone default (now() at time zone 'utc'),
  deleted_text boolean not null,
  url text not null,
  sent_text text not null
);

create table typer.messenger (
  unixt bigint not null,
  ts timestamp not null,
  deleted_text boolean not null,
  fbid text not null,
  message text not null
);

-- LNews

create table lnews.card (
  unixt bigint not null,
  ts timestamp not null,
  card text not null
);

create table lnews.panel (
  unixt bigint not null,
  ts timestamp not null,
  dismissed boolean not null default false,
  panel text not null
);

-- Bio_samples

create table bio.heartrate(
  start_unixt bigint not null,
  end_unixt bigint not null,
  measurement smallint not null
);

-- Emotions

create table emotion.evaluation(
  unixt bigint not null,
  ts timestamp not null,
  accomplished_eval smallint,
  social_eval smallint,
  exhausted_eval smallint,
  tired_eval smallint,
  happy_eval smallint,
  comments text
);

-- Events

-- The emotion columns describe how I felt afterwards
create table life_event.mark (
  unixt bigint not null,
  ts timestamp not null,

  anticipate boolean not null,
  start_unixt bigint not null,
  start_ts timestamp not null,
  event_unixt bigint not null,
  event_ts timestamp not null,
  end_unixt bigint not null,
  end_ts timestamp not null,

  anger boolean not null,
  contempt boolean not null,
  disgust boolean not null,
  fear boolean not null,
  interest boolean not null,
  joy boolean not null,
  sad boolean not null,
  surprise boolean not null,
  comment text
);

-- Lizzie Peaks

create table peaks.skill (
  time_learned_unixt bigint not null,
  time_learned_ts timestamp not null,
  time_spent_learning int not null,
  concept text not null,
  new_learnings text not null,
  old_skills text not null,
  percent_new smallint not null
);

create table peaks.review (
  time_learned_unixt bigint not null,
  time_learned_ts timestamp not null,
  time_reviewed_unixt bigint not null,
  time_reviewed_ts bigint not null,
  concept text not null,
  new_learnings text not null,
  time_spent_reviewing int not null
);

create table peaks.scheduled_review (
  time_learned_unixt bigint not null,
  time_learned_ts timestamp not null,
  time_scheduled_unixt bigint not null,
  time_scheduled_ts timestamp not null,
  concept text not null,
  scheduled_duration int not null
);