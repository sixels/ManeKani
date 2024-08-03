use time::{format_description::well_known::Rfc3339, macros::format_description, OffsetDateTime};

#[derive(Debug)]
pub struct HumanTime {
    pub rfc: String,
    pub short: String,
}

impl From<OffsetDateTime> for HumanTime {
    fn from(date: OffsetDateTime) -> Self {
        HumanTime {
            rfc: date.format(&Rfc3339).unwrap(),
            short: humanize_now(date),
        }
    }
}

fn humanize_now(date: OffsetDateTime) -> String {
    humanize_relative_date(date, OffsetDateTime::now_utc())
}

/// Returns how long ago the date was in human readable format.
fn humanize_relative_date(date: OffsetDateTime, since: OffsetDateTime) -> String {
    let diff = since - date;

    let hours = diff.whole_hours();

    match hours {
        0 => {
            let minutes = diff.whole_minutes();
            if minutes <= 10 {
                String::from("Just now")
            } else {
                format!("{} minutes ago", minutes)
            }
        }
        1..=23 => {
            if hours == 1 {
                String::from("An hour ago")
            } else {
                format!("{} hours ago", hours)
            }
        }
        24..=47 => String::from("Yesterday"),
        48..=167 => date
            .format(format_description!("[weekday repr:long]"))
            .unwrap(),
        _ => date
            .format(format_description!("[day] [month repr:short] [year]"))
            .unwrap(),
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_humanize_relative_date() {
        use time::{
            macros::{date, time},
            OffsetDateTime,
        };

        let date = OffsetDateTime::new_utc(date!(2024 - 01 - 01), time!(12:00:00));
        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 01), time!(12:00:00));

        assert_eq!(super::humanize_relative_date(date, since), "Just now");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 01), time!(12:11:00));
        assert_eq!(super::humanize_relative_date(date, since), "11 minutes ago");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 01), time!(13:00:00));
        assert_eq!(super::humanize_relative_date(date, since), "An hour ago");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 02), time!(12:00:00));
        assert_eq!(super::humanize_relative_date(date, since), "Yesterday");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 03), time!(12:00:00));
        assert_eq!(super::humanize_relative_date(date, since), "Monday");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 07), time!(12:00:00));
        assert_eq!(super::humanize_relative_date(date, since), "Monday");

        let since = OffsetDateTime::new_utc(date!(2024 - 01 - 10), time!(12:00:00));
        assert_eq!(
            super::humanize_relative_date(date, since),
            "01 Jan 2024"
        );
    }
}
