package moderation

var GuildModerationSchema = []string{`
CREATE TABLE IF NOT EXISTS moderation_config (
	guild_id TEXT PRIMARY KEY,
	mod_log TEXT,
	last_case_id BIGINT DEFAULT 0 NOT NULL
);
`,`
CREATE TABLE IF NOT EXISTS moderation_cases (
	case_id BIGINT NOT NULL,
	guild_id TEXT NOT NULL,
	staff_id TEXT NOT NULL,
	offender_id TEXT NOT NULL,
	reason TEXT,
	action TEXT NOT NULL,
	logLink TEXT NOT NULL,
    PRIMARY KEY (guild_id, case_id),
	CONSTRAINT fk_guild_cases FOREIGN KEY (guild_id)
		REFERENCES moderation_config (guild_id) ON DELETE CASCADE
);
`,
}
