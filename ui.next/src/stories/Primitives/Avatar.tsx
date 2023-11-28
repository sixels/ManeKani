import FallbackAvatar from "boring-avatars";

interface AvatarProps {
	/** The avatar image URL */
	url?: string;
	/** The username used to generate a fallback image if the given image is not available */
	username: string;
	/** The avatar image size */
	size?: "sm" | "md" | "lg"; // 40 60 80
	/** Should the avatar be square */
	isSquare?: boolean;
}

const getShapeClasses = (isSquare: boolean) => {
	return isSquare
		? "rounded-lg after:rounded-lg"
		: "rounded-full after:rounded-full";
};

export const Avatar = ({
	url,
	username,
	size = "md",
	isSquare = false,
}: AvatarProps) => {
	const shapeClasses = getShapeClasses(isSquare);

	const [imgSize, sizeClasses] = (
		{
			sm: [40, "w-[40px] h-[40px]"],
			md: [60, "w-[60px] h-[60px]"],
			lg: [80, "w-[80px] h-[80px]"],
		} as { [key: string]: [number, string] }
	)[size];

	return (
		<div
			className={`${shapeClasses} ${sizeClasses} overflow-hidden relative inline-flex after:w-full after:h-full after:contents-['""'] after:ring-inset after:ring-2 after:ring-gray-300/25 after:absolute after:left-0 after:top-0`}
		>
			{url ? (
				<img
					src={url}
					alt={`${username} avatar`}
					width={imgSize}
					height={imgSize}
					className="w-full h-full"
				/>
			) : (
				<DefaultAvatar username={username} size={imgSize} />
			)}
		</div>
	);
};

const DefaultAvatar = ({
	username,
	size,
}: {
	username: string;
	size: number;
}) => {
	const fallbackAvatarColors = [
		"#F08D1D",
		"#A764CD",
		"#E9ECEB",
		"#443AF8",
		"#F99FE3",
	];

	return (
		<FallbackAvatar
			name={username}
			size={size}
			variant="marble"
			colors={fallbackAvatarColors}
			square
		/>
	);
};
