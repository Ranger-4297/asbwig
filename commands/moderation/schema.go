package moderation

var GuildModerationSchema = []string{`
CREATE TABLE IF NOT EXISTS moderation_config (
	guild_id TEXT PRIMARY KEY,
	enabled BOOL DEFAULT FALSE NOT NULL,
	enabled_trigger_deletion BOOL DEFAULT FALSE NOT NULL,
	seconds_to_delete_trigger INT DEFAULT 0 NOT NULL,
	enabled_response_deletion BOOL DEFAULT FALSE NOT NULL,
	seconds_to_delete_response INT DEFAULT 0 NOT NULL,
	mod_log TEXT,
	manage_mute_role BOOL DEFAULT FALSE NOT NULL,
	mute_role TEXT,
	update_roles TEXT[],
	last_case_id BIGINT DEFAULT 0 NOT NULL
);
`,`
CREATE TABLE IF NOT EXISTS moderation_config_roles (
    guild_id TEXT NOT NULL,
    action_type TEXT NOT NULL,
    role_id TEXT NOT NULL,
    PRIMARY KEY (action_type, role_id),
    CONSTRAINT fk_moderation_config_roles_guild FOREIGN KEY (guild_id)
        REFERENCES moderation_config (guild_id) ON DELETE CASCADE
);
`,`
CREATE TABLE IF NOT EXISTS moderation_mutes (
    guild_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    roles TEXT[] NOT NULL,
    PRIMARY KEY (guild_id, user_id),
    CONSTRAINT fk_moderation_config_roles_guild FOREIGN KEY (guild_id)
        REFERENCES moderation_config (guild_id) ON DELETE CASCADE
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
