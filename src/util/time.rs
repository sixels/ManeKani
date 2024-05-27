use time::{Duration, OffsetDateTime};

#[derive(Debug)]
pub struct HumanTime {
    pub rfc: String,
    pub short: String,
}

impl From<OffsetDateTime> for HumanTime {
    fn from(date: OffsetDateTime) -> Self {
        HumanTime {
            rfc: date
                .format(time::macros::format_description!(
                "[weekday repr:short], [day] [month repr:short] [year] [hour]:[minute]:[second]"
            ))
                .unwrap(),
            short: humanize_date(date),
        }
    }
}

fn humanize_date(date: OffsetDateTime) -> String {
    // check if the date is today
    let now = OffsetDateTime::now_utc();
    if date.date() == now.date() {
        let hours_ago = now.hour() - date.hour();
        if hours_ago == 0 {
            String::from("Just now")
        } else {
            format!("{} hours ago", hours_ago)
        }
    }
    // check if the date was this week
    else if date.date() > now.date().saturating_sub(Duration::days(7)) {
        let days_ago = now.day() - date.day();
        if days_ago == 1 {
            String::from("Yesterday")
        } else {
            format!("{} days ago", days_ago)
        }
    } else {
        date.format(time::macros::format_description!(
            "[weekday repr:short], [day] [month repr:short] [year]"
        ))
        .unwrap()
    }
}
